package main

import (
	"encoding/json"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/splay"
	"log"
	"net/http"
)

var dataStore *splay.Tree

// KeyValueCouple models the input data for the POST request
type KeyValueCouple struct {
	Key   string    //`json:"key"`
	Value splay.Any //`json:"value"`
}

// APIError models the error object sent back to the client on error
type APIError struct {
	Code    int
	Message string
}

func main() {
	log.Println("Starting server...")
	log.Println("Initializing data store...")
	dataStore = splay.New()
	server := mux.NewRouter()
	server.HandleFunc("/api/v1/store/", handlePostValueByKey).
		Methods("POST")
	server.HandleFunc("/api/v1/store/{key}", handleGetValueByKey).
		Methods("GET")

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleGetValueByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(value)
}

func handlePostValueByKey(w http.ResponseWriter, r *http.Request) {
	var reqBody KeyValueCouple

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		apiErr := APIError{Message: err.Error()}
		if err := json.NewEncoder(w).Encode(apiErr); err != nil {
			panic(err)
		}
		return
	}

	if reqBody.Key == "" /*|| len(reqBody.Value) == 0*/ {
		http.Error(w, "Please specify a valid key and a value",
			http.StatusBadRequest)
		return
	}
	_, err := dataStore.Insert(reqBody.Key, reqBody.Value)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
