package controller

import (
	"PluginServer/config"
	"PluginServer/dto"
	"PluginServer/email"
	"PluginServer/ent"
	"PluginServer/hashing"
	"PluginServer/middleware"
	"PluginServer/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client) {
		userEndpoint := server.Group("user")
		userEndpoint.Static("/pfp", config.Config.CDN.Directory)

		userEndpoint.POST("", handleCreateUser)
		userEndpoint.GET("/:id", getUser)

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.PUT("", setPfp)
		userEndpoint.PATCH("", handleUpdateUser)
		userEndpoint.DELETE("", handleDeleteUser)
		userEndpoint.GET("/verify", verifyEmail)
		userEndpoint.GET("/resend", resendVerificationEmail)
	})
}

func getUser(ctx echo.Context) error {
	rawUserId := ctx.Param("id")

	userId, err := uuid.Parse(rawUserId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	userData, err := service.GetUser(userId)

	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
		}
		if err.Error() == "user deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{"message": "user deleted"})
		}

		log.Errorf("error querying user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"name":      userData.Name,
		"email":     userData.Email,
		"createdAt": userData.CreatedAt,
		"updatedAt": userData.UpdatedAt,
	})
}

func handleCreateUser(ctx echo.Context) error {

	userInfo := dto.CreateUserDTO{}
	ctx.Bind(&userInfo)

	// check if all the fields are provided

	_, err := mail.ParseAddress(userInfo.Email)
	isEmailValid := err == nil

	if strings.TrimSpace(userInfo.Name) == "" || strings.TrimSpace(userInfo.Password) == "" || !isEmailValid {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	userId, err := service.CreateUser(userInfo.Name, userInfo.Email, hashing.HashPassword(userInfo.Password))

	// error checking
	if err != nil {
		if ent.IsConstraintError(err) {
			if strings.Contains(err.Error(), "name") {
				return ctx.JSON(http.StatusConflict, echo.Map{
					"message": "username taken",
				})
			}
			return ctx.JSON(http.StatusConflict, echo.Map{
				"message": "email already in use",
			})
		}

		// log any uncaught errors
		log.Errorf("uncaught error creating user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	// start a session for email verification
	sessionId, err := service.StartVerifySession(*userId)
	fmt.Println(sessionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	// close session after 3 min
	time.AfterFunc(time.Minute*3, func() {
		service.CloseVerifySession(sessionId)
	})

	err = email.SendVerificationEmail(userInfo.Email, sessionId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"uuid": userId.String()})
}

func handleUpdateUser(ctx echo.Context) error {

	updateInfo := dto.UpdateUserDTO{}
	ctx.Bind(&updateInfo)

	// check if no fields are provided
	if strings.TrimSpace(updateInfo.Name) == "" && strings.TrimSpace(updateInfo.Email) == "" && strings.TrimSpace(updateInfo.Password) == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	if strings.TrimSpace(updateInfo.Password) != "" {
		updateInfo.Password = hashing.HashPassword(updateInfo.Password)
	}

	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))
	err := service.UpdateUser(updateInfo, userId)

	// error checking
	if err != nil {

		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})
		}
		if ent.IsValidationError(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": "validation error",
			})
		}
		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": "constraint error",
			})
		}
		if err.Error() == "user deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{
				"message": "user deleted",
			})
		}

		log.Errorf("uncaught error updating user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.NoContent(http.StatusOK)
}

func setPfp(ctx echo.Context) error {
	// get params from multipart
	userId := ctx.Request().Header.Get("UUID")
	file, err := ctx.FormFile("file")

	// validate request
	if err != nil || userId == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	userUUID, err := uuid.Parse(userId)
	if !service.DoesUserWithUUIDExist(userUUID) {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
	}

	// open file
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid file format"})
	}
	defer src.Close()

	// create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, userId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	// write to file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return ctx.NoContent(http.StatusOK)
}

func handleDeleteUser(ctx echo.Context) error {

	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	err := service.DeleteUser(userId)

	// error checking
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})
		}
		if err.Error() == "already deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{
				"message": "user already deleted",
			})
		}

		log.Errorf("uncaught error deleting user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.NoContent(http.StatusOK)
}

func verifyEmail(ctx echo.Context) error {

	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))
	sessionId, _ := uuid.Parse(ctx.QueryParam("sid"))

	err := service.CloseVerifySession(sessionId)

	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "session not found"})
		}

		log.Errorf("uncaught error closing verify session: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	err = service.VerifyUserEmail(userId)

	if err != nil {
		log.Errorf("uncaught error verifying user email: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.NoContent(200)
}

func resendVerificationEmail(ctx echo.Context) error {

	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))
	user, _ := service.GetUser(userId)

	if user.Verified {
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "user already verified"})
	}

	// close any existing sessions
	service.CloseVerifySessionByUserId(userId)
	sessionId, err := service.StartVerifySession(userId)

	if err != nil {
		log.Errorf("uncaught error starting verify session: %v", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	// close session after 3 min
	time.AfterFunc(time.Minute*3, func() {
		service.CloseVerifySession(sessionId)
	})

	err = email.SendVerificationEmail(user.Email, sessionId)

	if err != nil {
		log.Errorf("uncaught error sending user email: %v", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.NoContent(200)
}
