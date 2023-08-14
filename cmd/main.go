package main

import (
	"fmt"
	"go-server/internal/user"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"go-server/pkg/logger"
	"go-server/pkg/utils/router"
	"net/http"
)

func main() {
	cfg := config.GetConfig(".env")
	log := logger.GetLogger()
	db, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		log.Err(err.Error())
		panic(err)
	}
	ur := repositories.NewUserRepository(db)
	userService := user.NewUserService(ur)
	userCtrl := user.NewUserCtrl(userService)

	// log.Info("%+v\n", user)

	r := router.NewRouter()
	// routes
	r.GET("/users/:id", userCtrl.GetUserHandler)

	log.Info("Server started on port:%d", cfg.Server.Port)
	http.ListenAndServe(":"+fmt.Sprint(cfg.Server.Port), r.Handle())

}
