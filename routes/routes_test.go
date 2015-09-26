package routes

import (
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
	dataStore := splay.New()
	server = mux.NewRouter()
	handler := handlers.New(dataStore)
	SetupRoutes(server, handler)
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

	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"test\"}", key)
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

	bodyValue := fmt.Sprintf("{\"key\": \"%s\" \"value\": {\"Name\": \"Valerio\", \"Lastname\": \"Gheri\"}}", key)
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
	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"testValue\"}", key)
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
	bodyValue := fmt.Sprintf("{\"key\": \"%s\", \"value\": \"testValue\"}", key)
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

	getRespRec := httptest.NewRecorder()
	getReq, err := http.NewRequest("GET", "/api/v1/store/"+key, nil)
	if err != nil {
		t.Fatal("Creating 'GET /api/v1/store/'" + key + " request failed!")
	}

	server.ServeHTTP(getRespRec, getReq)
	if getRespRec.Code != http.StatusOK {
		t.Fatalf("Expected to receive status code %d, got %d",
			http.StatusOK, getRespRec.Code)
	}
}
