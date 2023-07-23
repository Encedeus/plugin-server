package service

import (
	"PluginServer/config"
	"PluginServer/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *ent.Client

func InitDB() *ent.Client {
	// Connect to database

	ctx := context.Background()

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

	// db.Schema.Create(context.Background())

	// update Db schema
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	Db = db
	return db
}
