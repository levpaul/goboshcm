package conn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	handler *mux.Router = ConstructRouter()
	connUrl string      = "http://localhost:5280/http-bind"
)

func TestOptionsEndpointExistsReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", connUrl, nil)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestOptionsEndpointReturnsNoBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", connUrl, nil)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, "", recorder.Body.String())
}

func TestOptionsEndpointReturnsMethods(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", connUrl, nil)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, "GET, POST, OPTIONS", recorder.Header().Get("Access-Control-Allow-Methods"))
}

func TestOptionsEndpointReturnsKleeneOrigin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", connUrl, nil)
	req.Header.Del("Origin")

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
}

func TestOptionsEndpointReturnsOriginalOrigin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", connUrl, nil)
	req.Header.Add("Origin", "https://mydomain.test.com")

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, "https://mydomain.test.com", recorder.Header().Get("Access-Control-Allow-Origin"))
}
