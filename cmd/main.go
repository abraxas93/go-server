package main

import (
	"go-server/internal/user"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"go-server/pkg/utils/router"
	"net/http"
)

func main() {
	cfg := config.GetConfig(".env")
	// log := logger.GetLogger()
	db, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		panic(err)
		return
	}
	ur := repositories.NewUserRepository(db)
	userService := user.NewUserService(ur)
	userCtrl := user.NewUserCtrl(userService)

	if err != nil {
		panic(err)
		return
	}
	// log.Info("%+v\n", user)

	r := router.NewRouter()

	r.GET("/users/:id", userCtrl.GetUserHandler)
	http.ListenAndServe(":8080", r.Handle())

}

//
