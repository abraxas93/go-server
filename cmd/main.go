package main

import (
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/logger"
)

func main() {
	cfg := config.GetConfig(".env")
	logger := logger.GetLogger()
	logger.Info("%+v\n", cfg)
	_, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		logger.Err(err.Error())
		panic(err)
	}

}
