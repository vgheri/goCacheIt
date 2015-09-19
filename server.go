package main

import (
	"encoding/json"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/splay"
	"log"
	"net/http"
)

const mimeTypeJSON string = "application/json; charset=UTF-8"

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
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if node == nil {
		writeJSONError(w, "Key not found", http.StatusNotFound)
		return
	}
	value, err := json.Marshal(node.Value)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", mimeTypeJSON)
	w.Write(value)
	return
}

func handlePostValueByKey(w http.ResponseWriter, r *http.Request) {
	var reqBody KeyValueCouple

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		writeJSONError(w, err.Error(), 422)
		return
	}

	if reqBody.Key == "" {
		writeJSONError(w, "Please specify a valid key and a value",
			http.StatusBadRequest)
		return
	}
	_, err := dataStore.Insert(reqBody.Key, reqBody.Value)
	if err != nil {
		writeJSONError(w, err.Error(),
			http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func writeJSONError(w http.ResponseWriter, errorMsg string, httpErrorCode int) {
	w.Header().Set("Content-Type", mimeTypeJSON)
	w.WriteHeader(httpErrorCode)
	apiErr := APIError{Message: errorMsg}
	if err := json.NewEncoder(w).Encode(apiErr); err != nil {
		panic(err)
	}
}
