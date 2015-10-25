package main

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/handlers"
	"github.com/vgheri/goCacheIt/routes"
	"github.com/vgheri/goCacheIt/splay"
	"log"
	"net/http"
)

const mimeTypeJSON string = "application/json; charset=UTF-8"

var dataStore *splay.Tree

func main() {
	log.Println("Starting server...")
	log.Println("Initializing data store...")
	// TODO read from flags
	var maxMemory uint64 = 1
	dataStore = splay.New(maxMemory)
	server := mux.NewRouter()
	handler := handlers.New(dataStore)
	routes.SetupRoutes(server, handler)
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
