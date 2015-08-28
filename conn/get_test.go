package conn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	getTestHandler *mux.Router = ConstructRouter()
	getTestConnUrl string      = "http://localhost:5280/http-bind"
)

func TestGetEndpointExistsReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", optionsTestConnUrl, nil)

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestOptionsEndpointReturnsBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", optionsTestConnUrl, nil)

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, `<html>
  <body>
    <a href='http://www.xmpp.org/extensions/xep-0124.html'>XEP-0124</a> - BOSH
  </body>
</html>`, recorder.Body.String())
}
