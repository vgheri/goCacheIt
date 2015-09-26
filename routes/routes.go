package routes

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/handlers"
)

// SetupRoutes setup routes for handling web commands
func SetupRoutes(router *mux.Router, handler *handlers.Handler) {
	router.HandleFunc("/api/v1/store/", handler.HandleAddValue).
		Methods("POST")
	router.HandleFunc("/api/v1/store/{key}", handler.HandleGetValue).
		Methods("GET")
}
