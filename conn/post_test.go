package conn

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	postTestHandler *mux.Router = ConstructRouter()
	postTestConnUrl string      = "http://localhost:5280/http-bind"
)

func TestOptionsParseXml(t *testing.T) {
	recorder := httptest.NewRecorder()

	var xmlStr = []byte(`<body rid='123456' xmlns='http://jabber.org/protocol/httpbind' to='chat.hello.com' />`)
	req, _ := http.NewRequest("POST", optionsTestConnUrl, bytes.NewBuffer(xmlStr))

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
