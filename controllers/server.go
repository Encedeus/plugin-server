package controllers

import (
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/ent"
	encMiddleware "github.com/Encedeus/pluginServer/middleware"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	registerRoutes(*Server)
}

func registerControllerRoutes(srv *Server, cs ...Controller) {
	for _, c := range cs {
		c.registerRoutes(srv)
	}
}

type Server struct {
	*echo.Echo
	DB *ent.Client
}

func NewEmptyServer(db *ent.Client) *Server {
	srv := &Server{
		Echo: echo.New(),
		DB:   db,
	}

	return srv
}

func WrapServerWithDefaults(srv *Server, _ *ent.Client) {
	srv.Use(encMiddleware.JSONSyntaxMiddleware) // json syntax checker
	srv.Use(encMiddleware.CORSMiddleware)       // cors config

	InitRouter(srv)
}

func NewDefaultServer(db *ent.Client) *Server {
	srv := NewEmptyServer(db)
	WrapServerWithDefaults(srv, db)

	return srv
}

func InitRouter(srv *Server) {
	registerControllerRoutes(srv,
		AuthController{},
		UserController{},
		PluginController{},
	)
}

func StartServer(srv *Server) {
	srv.Logger.Fatal(srv.Start(config.Config.Server.URI()))
}

func StartDefaultServer(db *ent.Client) {
	StartServer(NewDefaultServer(db))
}
