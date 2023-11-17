package db

import (
	"context"
	"fmt"
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/migrate"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var database *ent.Client

func InitDB() *ent.Client {
	// Connect to database

	db, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Config.DB.Host,
			config.Config.DB.Port,
			config.Config.DB.User,
			config.Config.DB.DBName,
			config.Config.DB.Password,
		),
	)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	err = db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	// update Db schema
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	database = db

	return db
}

func GetDb() *ent.Client {
	return database
}
