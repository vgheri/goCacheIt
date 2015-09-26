package handlers

import (
	"encoding/json"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/splay"
	//"log"
	"net/http"
)

// Handler is a strictly typed object containing the list of available handlers
type Handler struct {
	HandleGetValue func(w http.ResponseWriter, r *http.Request)
	HandleAddValue func(w http.ResponseWriter, r *http.Request)
}

const mimeTypeJSON string = "application/json; charset=UTF-8"

var dataStore *splay.Tree

// New initializes the package with the underlying data store instance
func New(store *splay.Tree) *Handler {
	dataStore = store
	handler := &Handler{
		HandleGetValue: getValue,
		HandleAddValue: addValue,
	}
	return handler
}

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

// getValue retrieves a value by key from the datastore
func getValue(w http.ResponseWriter, r *http.Request) {
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

// addValue adds a couple {key,value} to the datastore
func addValue(w http.ResponseWriter, r *http.Request) {
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
