package main

import (
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/controllers"
	"github.com/Encedeus/pluginServer/db"
	_ "github.com/Encedeus/pluginServer/ent/runtime"
)

func main() {
	config.InitConfig()
	db := db.InitDB()
	// go module.Init()
	controllers.StartDefaultServer(db)
}
