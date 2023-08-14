package router

import (
	"net/http"
	"strings"
)

// Types
type Request struct {
	*http.Request
	Params map[string]string
}
type HandlerFunc func(w http.ResponseWriter, r *Request)
type Route string
type Handlers map[Route]HandlerFunc
type HttpMethod string
type Routes map[HttpMethod]Handlers
type Router struct {
	routes Routes
}

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	req := &Request{
		Request: r,
		Params:  m,
	}
	fn(w, req)
}

// Constructor
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

// Private
func (r *Router) addRoute(method HttpMethod, pattern Route, handler HandlerFunc) {
	r.routes[method][pattern] = handler
}

func (r *Router) getRoutesByMethod(url string, method HttpMethod) []string {
	segmentsLength := len(strings.Split(url, "/"))
	handlers := r.routes[method]
	var routes []string
	for route := range handlers {
		if len(strings.Split(string(route), "/")) == segmentsLength && strings.Contains(string(route), "/:") {
			// check that it contains /:
			routes = append(routes, string(route))
		}
	}
	return routes
}

func parseSegments(url string) map[string]int {
	m := make(map[string]int)
	parts := strings.Split(url, "/")
	for i, s := range parts {
		m[s] = i
	}
	return m
}

func isUrlMatchRoute(url map[string]int, route map[string]int) bool {
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

func collectParams(url string, route string) map[string]string {
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

func (r *Router) baseHandler(w http.ResponseWriter, rq *Request) {

	method := rq.Method
	handlers := r.routes[HttpMethod(method)]
	handler := handlers[Route(rq.URL.Path)]

	if handler != nil {
		handler(w, rq)
		return
	}

	routes := r.getRoutesByMethod(rq.URL.Path, HttpMethod(method))

	segments := parseSegments(rq.URL.Path)

	for _, route := range routes {
		isMatching := isUrlMatchRoute(segments, parseSegments(route))
		if isMatching {
			params := collectParams(rq.URL.Path, route)
			handler = handlers[Route(route)]
			rq.Params = params
			handler(w, rq)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

// Public
func (r *Router) Handle() http.Handler {
	return HandlerFunc(r.baseHandler)
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
