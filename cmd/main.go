package main

import (
	"context"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"go-server/pkg/logger"
)

func main() {
	cfg := config.GetConfig(".env")
	log := logger.GetLogger()
	log.Info("%+v", cfg)
	db, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		log.Err(err.Error())
		panic(err)
	}
	ctx := context.Background()
	ur := repositories.New(db)
	user, err := ur.FindByID(ctx, 2)
	if err != nil {
		log.Err(err.Error())
		return
	}
	log.Info("%+v", user)
}
