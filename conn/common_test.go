package conn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	commonTestHandler *mux.Router
	commonTestConnUrl string = "http://localhost:5280/http-bind"
)

func init() {
	// Setup handler to use common headers "handler"
	commonTestHandler = mux.NewRouter()
	commonTestHandler.Methods("OPTIONS").Path("/http-bind").HandlerFunc(getCommonHeaders)
}

func TestCommonHeadersReturnsKleeneOrigin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", commonTestConnUrl, nil)
	req.Header.Del("Origin")

	commonTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
}

func TestCommonHeadersReturnsOriginalOrigin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", commonTestConnUrl, nil)
	req.Header.Add("Origin", "https://mydomain.test.com")

	commonTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "https://mydomain.test.com", recorder.Header().Get("Access-Control-Allow-Origin"))
}

func TestCommonHeadersReturnsHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", commonTestConnUrl, nil)

	commonTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "Content-Type", recorder.Header().Get("Access-Control-Allow-Headers"))
}

func TestOptionsEndpointReturnsCredentials(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", commonTestConnUrl, nil)

	commonTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "true", recorder.Header().Get("Access-Control-Allow-Credentials"))
}
