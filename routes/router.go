package routes

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/handlers"
	"net/http"
)

// NewRouter returns a router with all routes registered
func NewRouter(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := setupRoutes(handler)
	for _, route := range routes {
		var routeHandler http.Handler
		routeHandler = route.HandlerFunc
		routeHandler = middleware(routeHandler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(routeHandler)
	}

	return router
}
