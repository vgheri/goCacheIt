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
	fmt.Println("Starting server.")
	fmt.Printf("Max memory to be used: %d MB", maxMemory)
	dataStore := splay.New(maxMemory)
	fmt.Println("Initializing web server...")
	handler := handlers.New(dataStore)
	server := routes.NewRouter(handler)
	http.Handle("/", server)
	port := fmt.Sprintf(":%d", webServerPort)
	log.Printf("Server started and listening on port %d.", webServerPort)
	log.Fatal(http.ListenAndServe(port, nil))
}
