package main

import (
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/logger"
)

func main() {
	cfg := config.GetConfig(".env")
	log := logger.GetLogger()
	log.Info("%+v", cfg)
	_, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		log.Err(err.Error())
		panic(err)
	}
	log.Warn("Warning ...")
}
