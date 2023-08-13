package router

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Route string
type Handlers map[Route]HandlerFunc

type HttpMethod string
type Routes map[HttpMethod]Handlers

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

// func (r *Router) getRoutesByMethod(url string, method HttpMethod) []string {
// 	segmentsLength := len(strings.Split(url, "/"))
// 	handlers := r.routes[method]
// 	var routes []string
// 	for route := range handlers {
// 		if len(strings.Split(string(route), "/")) == segmentsLength && strings.Contains(string(route), "/:") {
// 			// check that it contains /:
// 			routes = append(routes, string(route))
// 		}
// 	}
// 	return routes
// }

func (r *Router) addRoute(method HttpMethod, pattern Route, handler HandlerFunc) {
	r.routes[method][pattern] = handler
}

func (r *Router) GET(pattern Route, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) POST(pattern Route, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) PUT(pattern Route, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) DELETE(pattern Route, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (router *Router) baseHandler(w http.ResponseWriter, r *http.Request) {
	// here should be core parsing logic
	method := r.Method
	handlers := router.routes[HttpMethod(method)]
	handler := handlers[Route(r.URL.Path)]
	if handler != nil {
		handler(w, r)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func (r *Router) Handle() http.Handler {
	return HandlerFunc(r.baseHandler)
}
