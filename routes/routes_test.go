package routes

import (
	// "encoding/json"
	"fmt"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/handlers"
	"github.com/vgheri/goCacheIt/splay"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var server *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func setupTest() {
	dataStore := splay.New(50)
	handler := handlers.New(dataStore)
	server = NewRouter(handler)
	respRec = httptest.NewRecorder()
}

func TestGetValueReturnsNotFound(t *testing.T) {
	setupTest()
	req, err = http.NewRequest("GET", "/api/v1/store/teststring", nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/store/teststring' request failed!")
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusNotFound {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusNotFound, respRec.Code)
	}
}

func TestGetValueReturnsBadRequest(t *testing.T) {
	setupTest()
	key := "abasbdsbcbADSAjkJIi29291929299232kadkasjdkajdaskjdkasjdkajsdkajk" +
		"kdsSjsdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSj" +
		"sdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjas" +
		"kjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka"

	req, err = http.NewRequest("GET", "/api/v1/store/"+key, nil)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusBadRequest {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusBadRequest, respRec.Code)
	}
}

func TestAddValueReturnsBadRequest(t *testing.T) {
	setupTest()
	key := "abasbdsbcbADSAjkJIi29291929299232kadkasjdkajdaskjdkasjdkajsdkajk" +
		"kdsSjsdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSj" +
		"sdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjas" +
		"kjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka"

	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"test\", \"duration\": 100000}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusBadRequest {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusBadRequest, respRec.Code)
	}
}

func TestAddValueReturnsBadRequestOnInvalidDuration(t *testing.T) {
	setupTest()
	key := "abasbdsbc"

	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"test\", \"duration\": 0}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusBadRequest {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusBadRequest, respRec.Code)
	}
}

func TestAddValueReturnsUnprocessableRequest(t *testing.T) {
	setupTest()
	key := "abasbdsbcbADSAjkJIi29291929299232kadkasjdkajdaskjdkasjdkajsdkajk" +
		"kdsSjsdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSj" +
		"sdjaskjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjas" +
		"kjdkasjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka" +
		"sjdkajsdkjaskdjaksjdkasjdkasjkdajskdjaksdjkasjdkasjkdjSjsdjaskjdkasjdka"

	bodyValue := fmt.Sprintf("{\"key\": \"%s\" \"value\": {\"Name\": \"Valerio\", \"Lastname\": \"Gheri\"}, \"duration\": 100000}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != 422 {
		t.Fatalf("Expected to receive status code %d, got %d",
			422, respRec.Code)
	}
}

func TestAddValueReturnsCreated(t *testing.T) {
	setupTest()
	key := "testKey"
	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"testValue\", \"duration\": 100000}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusCreated {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusCreated, respRec.Code)
	}
}

func TestGetValueReturnsValue(t *testing.T) {
	setupTest()
	key := "testKey"
	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"testValue\", \"duration\": 100000}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatal("Creating 'POST /api/v1/store/' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusCreated {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusCreated, respRec.Code)
	}

	getRespRec := httptest.NewRecorder()
	getReq, err := http.NewRequest("GET", "/api/v1/store/"+key, nil)
	if err != nil {
		t.Fatalf("Creating 'GET /api/v1/store/%s' request failed!", key)
	}

	server.ServeHTTP(getRespRec, getReq)
	if getRespRec.Code != http.StatusOK {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusOK, getRespRec.Code)
	}
}

func TestRemoveValueShouldReturnNotFoundWhenKeyDoesntExist(t *testing.T) {
	setupTest()
	key := "TestRemoveValueShouldReturnNotFoundWhenKeyDoesntExist"
	req, err = http.NewRequest("DELETE", "/api/v1/store/"+key, nil)
	if err != nil {
		t.Fatalf("Creating 'DELETE /api/v1/store/%s' request failed!", key)
	}
	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusNotFound {
		t.Fatalf("Expected to receive status code %d, got %d instead.",
			http.StatusNotFound, respRec.Code)
	}
}

func TestRemoveValueShouldReturnNoContentWhenAValueIsSuccessfullyRemoved(t *testing.T) {
	setupTest()
	key := "TestRemoveValueShouldReturnNoContentWhenAValueIsSuccessfullyRemoved"
	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"testValue\", \"duration\": 100000}", key)
	body := strings.NewReader(bodyValue)
	req, err = http.NewRequest("POST", "/api/v1/store/", body)
	if err != nil {
		t.Fatal("Creating 'POST /api/v1/store/' request failed!", key)
	}

	server.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusCreated {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusCreated, respRec.Code)
	}

	delRespRec := httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/api/v1/store/"+key, nil)
	if err != nil {
		t.Fatalf("Creating 'DELETE /api/v1/store/%s' request failed!", key)
	}
	server.ServeHTTP(delRespRec, req)
	if delRespRec.Code != http.StatusNoContent {
		t.Fatalf("Expected to receive status code %d, got %d instead.",
			http.StatusNoContent, delRespRec.Code)
	}
}
