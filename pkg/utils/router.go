package router

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handlers map[string]func(http.ResponseWriter, *http.Request)

type Routes map[string]Handlers

type Router struct {
	routes Routes
}

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn(w, r)
}

func NewRouter() Router {
	return Router{
		routes: Routes{
			"GET":    Handlers{},
			"POST":   Handlers{},
			"PUT":    Handlers{},
			"DELETE": Handlers{},
		},
	}
}

// should store paths patterns
// first defining map for storing method types
// should parse url
// should store handler to corresponding pattern
