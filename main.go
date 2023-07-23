package main

import (
	"PluginServer/config"
	"PluginServer/server"
	"PluginServer/service"
)

func main() {
	config.InitConfig()
	service.InitDB()
	server.ServerInit()
}
