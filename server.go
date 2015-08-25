package main

import (
	"fmt"
	//"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// server := mux.NewRouter()
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside the handler")
		fmt.Fprintf(w, "Hello!")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
