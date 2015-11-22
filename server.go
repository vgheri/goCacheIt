package main

import (
	"fmt"
	"github.com/vgheri/goCacheIt/handlers"
	"github.com/vgheri/goCacheIt/routes"
	"github.com/vgheri/goCacheIt/splay"
	"log"
	"net/http"
)

const mimeTypeJSON string = "application/json; charset=UTF-8"

func main() {
	log.Println("Starting server...")
	log.Printf("Max memory to be used: %d MB", maxMemory)
	log.Println("Initializing data store...")
	dataStore := splay.New(maxMemory)
	log.Println("Initializing web server...")
	handler := handlers.New(dataStore)
	server := routes.NewRouter(handler)
	http.Handle("/", server)
	port := fmt.Sprintf(":%d", webServerPort)
	log.Printf("Server listening on port %d...", webServerPort)
	log.Fatal(http.ListenAndServe(port, nil))
}
