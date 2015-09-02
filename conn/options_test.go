package conn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	optionsTestHandler *mux.Router = ConstructRouter()
	optionsTestConnUrl string      = "http://localhost:5280/http-bind"
)

func TestOPTIONSEndpointExistsReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", optionsTestConnUrl, nil)

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestOPTIONSEndpointReturnsNoBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", optionsTestConnUrl, nil)

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "", recorder.Body.String())
}

func TestOPTIONSEndpointReturnsMethods(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", optionsTestConnUrl, nil)

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, "GET, POST, OPTIONS", recorder.Header().Get("Access-Control-Allow-Methods"))
}
