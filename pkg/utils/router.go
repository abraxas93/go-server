package router

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handlers map[string]HandlerFunc

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

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	r.routes[method][pattern] = handler
}

func (r *Router) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) POST(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) PUT(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) DELETE(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

// should store paths patterns
// first defining map for storing method types
// should parse url
// should store handler to corresponding pattern
