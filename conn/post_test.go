package conn

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/xml"

	"github.com/gorilla/mux"
	"github.com/levilovelock/goboshcm/sessions"
	"github.com/stretchr/testify/assert"
)

var (
	postTestHandler *mux.Router = ConstructRouter()
	postTestConnUrl string      = "http://localhost:5280/http-bind"

	validFirstPOSTSessionCreationRequestBody = []byte(`<body
        rid='2902921866'
        xmlns='http://jabber.org/protocol/httpbind'
        to='chat.mysite.com'
        xml:lang='en'
        wait='60'
        hold='1'
        content='text/xml; charset=utf-8'
        ver='1.6'
        xmpp:version='1.0'
        xmlns:xmpp='urn:xmpp:xbosh'
        route='xmpp:mysite.com:5999'/>`)
)

func TestSessionCreationPOSTReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()
	sessions.InitialiseSessionsPool()

	req, _ := http.NewRequest("POST", postTestConnUrl, bytes.NewBuffer(validFirstPOSTSessionCreationRequestBody))

	postTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestPOSTWithBogusXmlReturns400(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", postTestConnUrl, bytes.NewBuffer([]byte(`<xml man i am cool>`)))

	postTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestParsingSessionCreationPayload(t *testing.T) {
	var pl Payload
	err := xml.Unmarshal(validFirstPOSTSessionCreationRequestBody, &pl)

	assert.Nil(t, err)

	assert.Equal(t, "2902921866", pl.RID)
	assert.Equal(t, "http://jabber.org/protocol/httpbind", pl.XMLNS)
	assert.Equal(t, "chat.mysite.com", pl.To)
	assert.Equal(t, "en", pl.XMLLang)
	assert.Equal(t, "60", pl.Wait)
	assert.Equal(t, "1", pl.Hold)
	assert.Equal(t, "text/xml; charset=utf-8", pl.Content)
	assert.Equal(t, "1.6", pl.Version)
	assert.Equal(t, "1.0", pl.XMPPVersion)
	assert.Equal(t, "urn:xmpp:xbosh", pl.XMLNSXMPP)
	assert.Equal(t, "xmpp:mysite.com:5999", pl.Route)
	assert.Equal(t, "", pl.SID)
}

func TestParsingSid(t *testing.T) {
	var pl Payload
	err := xml.Unmarshal([]byte(`<body sid="abc123" />`), &pl)

	assert.Nil(t, err)
	assert.Equal(t, "abc123", pl.SID)
}

func TestPOSTWithNonExistantSidReturns400(t *testing.T) {
	recorder := httptest.NewRecorder()
	sessions.InitialiseSessionsPool()

	req, _ := http.NewRequest("POST", postTestConnUrl, bytes.NewBuffer([]byte(`<body sid='a1b2c3' />`)))

	postTestHandler.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}
