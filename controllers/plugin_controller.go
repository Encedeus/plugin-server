package controllers

import (
	"fmt"
	"github.com/Encedeus/pluginServer/api"
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/middleware"
	"github.com/Encedeus/pluginServer/proto"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

type PluginController struct {
	Controller
}

func (pc PluginController) registerRoutes(srv *Server) {
	pluginEndpoint := srv.Group("plugin")
	{

		pluginEndpoint.GET("/:name", func(c echo.Context) error {
			return pc.HandleFindPlugin(c, srv.DB)
		})

		pluginEndpoint.GET("/readme/:id", func(c echo.Context) error {
			return pc.HandleGetReadme(c, srv.DB)
		})

		pluginEndpoint.GET("/search/", func(c echo.Context) error {
			return pc.HandlePluginSearch(c, srv.DB)
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
			releaseEndpoint.PUT("", func(c echo.Context) error {
				return pc.HandleDeprecateRelease(c, srv.DB)
			})
		}
	}
}

func (PluginController) HandleFindPlugin(c echo.Context, db *ent.Client) error {
	pluginIdentifier := c.Param("name")

	id, err := uuid.Parse(pluginIdentifier)

	var resp *protoapi.Plugin

	if err == nil {
		resp, err = services.FindPluginById(c.Request().Context(), db, id)
		if err != nil {
			return errors2.GetHTTPErrorResponse(c, err)
		}
	} else {
		resp, err = services.FindPluginByName(c.Request().Context(), db, pluginIdentifier)
		if err != nil {
			return errors2.GetHTTPErrorResponse(c, err)
		}
	}
	return proto.MarshalControllerProtoResponseToJSON(&c, 200, resp)
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

func (PluginController) HandlePublishRelease(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	publishReq := new(protoapi.PluginPublishReleaseRequest)
	c.Bind(publishReq)

	pluginId, err := uuid.Parse(publishReq.PluginId)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, errors2.ErrInvalidPluginId)
	}

	// check if plugin exists before calls to the GitHub api
	plugin, err := services.GetOwnedPluginWithSource(ctx, db, userId, pluginId)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	err = services.PublishRelease(ctx, db, publishReq, plugin)
	fmt.Println(err)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return c.NoContent(200)
}

func (PluginController) HandleDeprecateRelease(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IdFromAccessContext(ctx)

	deprecateReq := new(protoapi.PluginDeprecateReleaseRequest)
	c.Bind(deprecateReq)

	pluginId, err := uuid.Parse(deprecateReq.PluginId)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, errors2.ErrInvalidPluginId)
	}

	pluginData, err := services.GetPluginWithAllEdges(ctx, db, pluginId)
	if pluginData.Edges.Owner.ID != userId {
		return errors2.GetHTTPErrorResponse(c, errors2.ErrUnauthorized)
	}

	err = services.DeprecateRelease(ctx, db, pluginId, deprecateReq.ReleaseName)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return c.NoContent(200)
}

func (PluginController) HandleGetReadme(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	pluginId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, errors2.ErrInvalidPluginId)
	}

	publicationData, err := services.GetLatestPublication(ctx, db, pluginId)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	pluginData, err := services.GetPluginWithSource(ctx, db, pluginId)

	repo := proto.GithubUriToProtoGithubRepo(pluginData.Edges.Source.Repository)

	readme, err := api.GetReadme(repo, publicationData.Tag)

	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, 200, &protoapi.PluginGetReadmeResponse{Readme: readme})
}

func (PluginController) HandlePluginSearch(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	searchQuery := c.QueryParam("q")

	// request query parsing

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	pluginsPerPage, err := strconv.Atoi(c.QueryParam("perpage"))
	if err != nil {
		pluginsPerPage = 20
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 1000
	}

	if limit > 1000 {
		limit = 1000
	}

	// query db
	plugins, err := services.SearchPluginsByName(ctx, db, searchQuery, limit)
	if err != nil {
		return errors2.GetHTTPErrorResponse(c, err)
	}

	// handle slicing
	endIndex := page * pluginsPerPage
	startIndex := endIndex - pluginsPerPage

	if startIndex > len(plugins) {
		return errors2.GetHTTPErrorResponse(c, errors2.ErrInvalidPageNumber)
	}

	if endIndex > len(plugins) {
		endIndex = len(plugins)
	}

	pages := math.Ceil(float64(len(plugins)) / float64(pluginsPerPage))
	plugins = plugins[startIndex:endIndex]

	return proto.MarshalControllerProtoResponseToJSON(&c, 200, &protoapi.PluginSearchByNameResponse{
		Plugins: proto.EntPluginEntitiesToProtoPlugin(plugins),
		Pages:   int32(pages),
		Page:    int32(page),
	})
}
