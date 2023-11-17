package controllers

import (
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/middleware"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PluginController struct {
	Controller
}

func (pc PluginController) registerRoutes(srv *Server) {
	pluginEndpoint := srv.Group("plugin")
	{
		pluginEndpoint.Use(middleware.AccessJWTAuth)

		pluginEndpoint.POST("", func(c echo.Context) error {
			return pc.HandleCreatePlugin(c, srv.DB)
		})
	}
}

func (PluginController) HandleCreatePlugin(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	createReq := new(protoapi.PluginCreateRequest)

	err := c.Bind(createReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}
	_, err = services.CreatePlugin(ctx, db, createReq, userId)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return c.NoContent(200)
}
