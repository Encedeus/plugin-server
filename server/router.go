package server

import (
	"PluginServer/controller"
	"PluginServer/ent"
	"PluginServer/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteInit(server *echo.Echo, db *ent.Client) {
	server.Use(middleware.JSONSyntaxMiddleware)

	controller.InitControllers(server, db)

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
