package controller

import (
	"PluginServer/dto"
	"PluginServer/ent"
	"PluginServer/middleware"
	"PluginServer/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strings"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client) {
		pluginEndpoint := server.Group("plugin")
		{
			pluginEndpoint.GET("/:name", handleGetPlugin)

			pluginEndpoint.Use(middleware.AccessJWTAuth)

			pluginEndpoint.POST("", handleCreatePlugin)
			pluginEndpoint.PATCH("", handleUpdatePlugin)
		}
	})
}

func handleGetPlugin(ctx echo.Context) error {
	pluginName := ctx.Param("name")

	pluginData, err := service.GetPluginByName(pluginName)

	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(404, echo.Map{
				"message": "plugin does not exist",
			})
		}

		log.Errorf("uncaught error querying plugin: %v", err)

		return ctx.JSON(500, echo.Map{"message": "internal server error"})
	}

	return ctx.JSON(200, pluginData)

}

func handleCreatePlugin(ctx echo.Context) error {

	pluginInfo := dto.CreatePluginDTO{}
	ctx.Bind(&pluginInfo)

	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	if strings.TrimSpace(pluginInfo.Name) == "" {
		return ctx.JSON(400, echo.Map{
			"message": "bad request",
		})
	}

	err := service.CreatePlugin(pluginInfo, userId)

	if err != nil {

		if ent.IsConstraintError(err) {
			return ctx.JSON(409, echo.Map{
				"message": "plugin of same name already exists",
			})
		}

		log.Errorf("uncaught error creating plugin: %v", err)

		return ctx.JSON(500, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.NoContent(200)
}

func handleUpdatePlugin(ctx echo.Context) error {
	updateInfo := dto.UpdatePluginDTO{}
	ctx.Bind(&updateInfo)

	// TODO: finish this shit

	return nil
}
