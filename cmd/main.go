package main

import (
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/logger"
)

func main() {
	cfg := config.GetConfig(".env")
	logger.InitLogger(make(map[string]string))
	logger, _ := logger.GetLogger()
	logger.Info("%+v\n", cfg)
	postgres.Connect()
}
