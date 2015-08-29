package conn

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/xml"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	postTestHandler *mux.Router = ConstructRouter()
	postTestConnUrl string      = "http://localhost:5280/http-bind"

	validFirstPostRequestBody = []byte(`<body
        rid='2902921866'
        xmlns='http://jabber.org/protocol/httpbind'
        to='chat.mysite.com'
        xmllang='en'
        wait='60'
        hold='1'
        content='text/xml; charset=utf-8'
        ver='1.6'
        xmppversion='1.0'
        xmlnsxmpp='urn:xmpp:xbosh'
        route='xmpp:mysite.com:5999'/>`)
)

func TestPostReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", optionsTestConnUrl, bytes.NewBuffer(validFirstPostRequestBody))

	optionsTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestParsingPayload(t *testing.T) {
	var pl Payload
	err := xml.Unmarshal(validFirstPostRequestBody, &pl)

	assert.Nil(t, err)

	assert.Equal(t, 2902921866, pl.Rid)
	assert.Equal(t, "http://jabber.org/protocol/httpbind", pl.Xmlns)
	assert.Equal(t, "chat.mysite.com", pl.To)
	assert.Equal(t, "en", pl.XmlLang)
	assert.Equal(t, 60, pl.Wait)
	assert.Equal(t, 1, pl.Hold)
	assert.Equal(t, "text/xml; charset=utf-8", pl.Content)
	assert.Equal(t, "1.6", pl.Ver)
	assert.Equal(t, "1.0", pl.XmppVer)
	assert.Equal(t, "urn:xmpp:xbosh", pl.XmlnsXmpp)
	assert.Equal(t, "xmpp:mysite.com:5999", pl.Route)
}
