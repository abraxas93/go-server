package main

import (
	"fmt"
	"go-server/internal/user"
	"go-server/pkg/config"
	"go-server/pkg/database/postgres"
	"go-server/pkg/database/repositories"
	"go-server/pkg/utils/router"
	"net/http"
	"strings"
)

func ParseUrl(url string) []string {
	l := len(strings.Split("/users/43/comments", "/"))
	m := map[string]int{
		"/users/:id":             2,
		"/users/create":          3,
		"/users/:id/comments":    4,
		"/users/latest/comments": 4,
		"/users/admin/:id":       4,
	}
	var routes []string
	for key := range m {
		if m[key] == l && strings.Contains(key, "/:") {
			// check that it contains /:
			routes = append(routes, key)
		}
	}
	return routes
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL)
	// split by ?
	// match by segments
	// if there is no strict match then check params avialability starting from :id. If there is a match by params. Then compare other segments that are strict
	// first define which patterns matching url
	//
	//   /users/32/comments
	//   /users/32/posts
	//   /users/create/comment
	//   /users/:id/comments
	w.Write([]byte("Hello, world!"))
}

type Segment struct {
	Substring string
	Position  int
}

func ParseSegments(url string) map[string]int {
	m := make(map[string]int)
	parts := strings.Split(url, "/")
	for i, s := range parts {
		m[s] = i
	}
	return m
}

func CollectParams(url string, route string) map[string]string {
	params := make(map[string]string)
	urlParts := strings.Split(url, "/")
	routeParts := strings.Split(route, "/")
	for i, s := range routeParts {
		if !strings.Contains(s, ":") {
			continue
		}
		params[s] = urlParts[i]
	}
	return params
}

func IsUrlMatchRoute(url map[string]int, route map[string]int) bool {
	for key := range route {
		if strings.Contains(key, ":") {
			continue
		}
		if route[key] != url[key] {
			return false
		}
	}
	return true
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

	if err != nil {
		panic(err)
		return
	}
	// log.Info("%+v\n", user)

	r := router.NewRouter()

	r.GET("/users/:id", userCtrl.GetUserHandler)

	// Start the server
	http.ListenAndServe(":8080", r.Handle())

}

//
