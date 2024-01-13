package controllers

import (
	"fmt"
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/middleware"
	"github.com/Encedeus/pluginServer/proto"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type UserController struct {
	Controller
}

func (uc UserController) registerRoutes(srv *Server) {
	userEndpoint := srv.Group("user")
	{
		userEndpoint.Static("/pfp", config.Config.Storage.Directory)

		userEndpoint.GET("/:id", func(c echo.Context) error {
			return handleFindUser(c, srv.DB)
		})

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.GET("", func(c echo.Context) error {
			return handleGetSelf(c, srv.DB)
		})

		userEndpoint.PUT("", func(c echo.Context) error {
			return handleSetPfp(c, srv.DB)
		})
		userEndpoint.PATCH("", func(c echo.Context) error {
			return handleUpdateUser(c, srv.DB)
		})
		userEndpoint.DELETE("", func(c echo.Context) error {
			return handleDeleteUser(c, srv.DB)
		})
	}
}

func handleFindUser(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	rawUserId := c.Param("id")

	userId, err := uuid.Parse(rawUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid user id"})
	}

	resp, err := services.FindOneUser(ctx, db, userId)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func handleUpdateUser(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	updateReq := new(protoapi.UserUpdateRequest)
	err := c.Bind(updateReq)

	userData, err := services.UpdateUser(ctx, db, userId, updateReq)

	// error handling
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	resp := proto.EntUserEntityToProtoUser(userData)

	return proto.MarshalControllerProtoResponseToJSON(&c, 200, resp)
}

func handleSetPfp(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	// get params from multipart
	userId, _ := middleware.IdFromAccessContext(ctx)
	file, err := c.FormFile("file")

	// validate request
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	// open file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid file format"})
	}
	defer src.Close()

	// create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.Storage.Directory, userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	// write to file
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return c.NoContent(http.StatusOK)
}

func handleDeleteUser(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	authUUID, _ := middleware.IdFromAccessContext(ctx)

	err := services.DeleteUser(ctx, db, authUUID)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func handleGetSelf(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	userData, err := services.GetUser(ctx, db, userId)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, 200, proto.EntUserEntityToProtoUser(userData))
}
