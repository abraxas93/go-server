package main

import (
	"fmt"
	"go-server/internal/user"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn(w, r)
}

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
	user, err := userService.GetUser(2)
	if err != nil {
		panic(err)
		return
	}
	// log.Info("%+v\n", user)
	p := user
	fmt.Printf("%p\n", user)
	fmt.Printf("%p\n", p)
	mux := http.NewServeMux()

	mux.HandleFunc("/users/:id", userCtrl.HelloHandler)
	mux.HandleFunc("/users/", userCtrl.GetUserHandler)

	http.Handle("/", HandlerFunc(HelloHandler))
	// Start the server
	http.ListenAndServe(":8080", nil)

}

//
