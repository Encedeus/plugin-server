package controllers

import (
	"fmt"
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/hashing"
	"github.com/Encedeus/pluginServer/middleware"
	"github.com/Encedeus/pluginServer/proto"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

type AuthController struct {
	Controller
}

func (ac AuthController) registerRoutes(srv *Server) {
	authEndpoint := srv.Group("auth")
	{
		authEndpoint.POST("/register", func(c echo.Context) error {
			return ac.handleRegisterUser(c, srv.DB)
		})
		authEndpoint.POST("/signin", func(c echo.Context) error {
			return ac.handleUserSignIn(c, srv.DB)
		})

		authEndpoint.Use(middleware.RefreshJWTAuth)

		authEndpoint.GET("/refresh", func(c echo.Context) error {
			return ac.handleRefreshToken(c, srv.DB)
		})
		authEndpoint.DELETE("/signout", func(c echo.Context) error {
			return ac.handleSignOut(c, srv.DB)
		})
	}
}

func (AuthController) handleRegisterUser(c echo.Context, db *ent.Client) error {

	ctx := c.Request().Context()
	createReq := new(protoapi.UserRegisterRequest)
	err := c.Bind(createReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	user, err := services.CreateUser(ctx, db, createReq)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	accessToken, refreshToken, err := services.GetTokenPair(user.ID)

	// set refresh token cookie
	c.SetCookie(&http.Cookie{
		Name:     "encedeus_plugins_refreshToken",
		Value:    refreshToken,
		Secure:   true,
		Expires:  time.Now().Add(services.RefreshTokenExpireTime),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	return proto.MarshalControllerProtoResponseToJSON(&c, 200,
		&protoapi.UserAuthorizeResponse{
			AccessToken: accessToken,
		},
	)
}

func (AuthController) handleUserSignIn(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	signInReq := new(protoapi.UserSignInRequest)
	// error safe because of the json syntax middleware
	err := c.Bind(signInReq)

	var (
		passwordHash string
		userId       *uuid.UUID
	)

	// check which method was used for log in

	fmt.Println(signInReq.Uid, signInReq.Password)

	if _, err2 := mail.ParseAddress(signInReq.Uid); err2 == nil {
		passwordHash, userId, err = services.GetUserAuthDataAndHashByEmail(ctx, db, signInReq.Uid)
	} else if strings.TrimSpace(signInReq.Uid) != "" {
		passwordHash, userId, err = services.GetUserUUIDAndHashByUsername(ctx, db, signInReq.Uid)
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "either username or email must be specified",
		})
	}
	// handle errors
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	if !hashing.VerifyHash(signInReq.Password, passwordHash) {
		return c.JSON(401, echo.Map{
			"message": "unauthorized",
		})
	}

	// generate access and refresh tokens
	accessToken, refreshToken, err := services.GetTokenPair(*userId)

	// set refresh token cookie
	c.SetCookie(&http.Cookie{
		Name:     "encedeus_plugins_refreshToken",
		Value:    refreshToken,
		Secure:   true,
		Expires:  time.Now().Add(services.RefreshTokenExpireTime),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	})

	return proto.MarshalControllerProtoResponseToJSON(&c, 200,
		&protoapi.UserAuthorizeResponse{
			AccessToken: accessToken,
		},
	)
}

func (AuthController) handleRefreshToken(c echo.Context, _ *ent.Client) error {
	// error safe because of the RefreshJWTAuth middleware
	token, _ := services.GetRefreshTokenFromCookie(c)
	_, tokenData, _ := services.ValidateRefreshJWT(token)

	userId, _ := uuid.Parse(tokenData.ID)

	// generate access token
	accessToken, _ := services.GenerateAccessToken(userId)

	return proto.MarshalControllerProtoResponseToJSON(&c, 200,
		&protoapi.UserAuthorizeResponse{
			AccessToken: accessToken,
		},
	)
}

func (AuthController) handleSignOut(c echo.Context, _ *ent.Client) error {
	c.SetCookie(&http.Cookie{
		Name:     "encedeus_plugins_refreshToken",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.UnixMilli(0),
		Secure:   true,
		Path:     "/",
	})

	return c.NoContent(http.StatusOK)
}