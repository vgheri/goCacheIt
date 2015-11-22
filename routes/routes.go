package routes

import (
	"github.com/vgheri/goCacheIt/handlers"
	"net/http"
)

// Route maps key information for an HTTP route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of routes
type Routes []Route

func setupRoutes(handler *handlers.Handler) Routes {
	return Routes{
		Route{
			Name:        "GetValue",
			Method:      "GET",
			Pattern:     "/api/v1/store/{key}",
			HandlerFunc: handler.HandleGetValue,
		},
		Route{
			Name:        "AddValue",
			Method:      "POST",
			Pattern:     "/api/v1/store/",
			HandlerFunc: handler.HandleAddValue,
		},
	}
}
