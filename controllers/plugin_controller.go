package controllers

import (
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/middleware"
	"github.com/Encedeus/pluginServer/proto"
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
		pluginEndpoint.GET("/:id", func(c echo.Context) error {
			return pc.HandleFindPlugin(c, srv.DB)
		})

		pluginEndpoint.Use(middleware.AccessJWTAuth)

		pluginEndpoint.POST("", func(c echo.Context) error {
			return pc.HandleCreatePlugin(c, srv.DB)
		})

		releaseEndpoint := pluginEndpoint.Group("/release")
		{
			releaseEndpoint.POST("", func(c echo.Context) error {
				return pc.HandlePublishRelease(c, srv.DB)
			})
		}
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

func (PluginController) HandleFindPlugin(c echo.Context, db *ent.Client) error {
	pluginName := c.Param("id")

	resp, err := services.FindPluginByName(c.Request().Context(), db, pluginName)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, 200, resp)
}

func (PluginController) HandlePublishRelease(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	publishReq := new(protoapi.PluginPublishReleaseRequest)
	c.Bind(publishReq)

	// check if plugin exists before calls to the GitHub api
	plugin, err := services.GetPluginForReleasePublication(ctx, db, userId, publishReq.PluginName)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}
	// todo: remove this giant piece of sranje that is the next line

	err = services.PublishRelease(ctx, db, publishReq, plugin)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return c.NoContent(200)
}
