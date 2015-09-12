package main

import (
	"encoding/json"
	"fmt"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/splay"
	"log"
	"net/http"
)

var dataStore *splay.Tree

func main() {
	log.Println("Starting server...")
	log.Println("Initializing data store...")
	dataStore = splay.New()
	server := mux.NewRouter()
	server.HandleFunc("/test/{id}", handleTest).
		Methods("GET")
	server.HandleFunc("/api/v1/store/{key}", handleGetValueByKey).
		Methods("GET")

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Inside the handler. Id %s", id)
	fmt.Fprintf(w, "Hello id %s!", id)
}

func handleGetValueByKey(w http.ResponseWriter, r *http.Request) {
	//log.Print("Inside handleGetValueByKey.")
	vars := mux.Vars(r)
	key := vars["key"]
	//log.Printf("Key %s", key)
	node, err := dataStore.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if node == nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	value, err := json.Marshal(node.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(value)
}
