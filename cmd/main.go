package main

import (
	"fmt"
	"go-server/internal/user"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	cfg := config.GetConfig(".env")
	// log := logger.GetLogger()
	db, err := postgres.Connect(cfg.DB.Postgres)
	if err != nil {
		panic(err)
		return
	}
	ur := repositories.New(db)
	userService := user.NewService(ur)
	user, err := userService.GetUser(2)
	if err != nil {
		panic(err)
		return
	}
	// log.Info("%+v\n", user)
	p := user
	fmt.Printf("%p\n", user)
	fmt.Printf("%p\n", p)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", helloHandler)
	// // mux.HandleFunc("/user", user.UserCtrl.GetUserHandler)
	// // Start the server
	// http.ListenAndServe(":8080", mux)

}
