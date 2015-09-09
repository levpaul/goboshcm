package sessions

import (
	"encoding/xml"
	"strconv"
	"testing"

	"github.com/levilovelock/goboshcm/common"
	"github.com/stretchr/testify/assert"
)

// ========================================
// Helper Functions
// ========================================

func getSessionCreationPayload() *common.Payload {
	pl := &common.Payload{
		RID:         "2902921866",
		To:          "chat.mysite.com",
		XMLNS:       "http://jabber.org/protocol/httpbind",
		XMLLang:     "en",
		Wait:        "60",
		Hold:        "1",
		Content:     "text/xml; charset=utf-8",
		Version:     "1.6",
		XMPPVersion: "1.0",
		XMLNSXMPP:   "urn:xmpp:xbosh",
		Route:       "xmpp:mysite.com:5999",
	}
	return pl
}

type responsePayload struct {
	XMLName  xml.Name `xml:"body"`
	XMLNS    string   `xml:"xmlns,attr"`
	SID      string   `xml:"sid,attr"`
	Wait     int      `xml:"wait,attr"`
	Requests int      `xml:"requests,attr"`
	Polling  string   `xml:"polling,attr"`
}

func TestSessionCreationAndValidation(t *testing.T) {
	InitialiseSessionsPool()
	sid, err := CreateNewSession()
	assert.Nil(t, err)

	assert.True(t, SessionExists(sid))
}

func TestSidCreationReturnsSid40CharsLong(t *testing.T) {
	InitialiseSessionsPool()
	sid, _ := CreateNewSession()

	assert.Equal(t, 40, len(sid))
}

func TestGenerateSessionCreationResponseIsXmlAndHasMinimalAttrs(t *testing.T) {
	pl := getSessionCreationPayload()
	pl.SID, _ = CreateNewSession()

	stringResponse, err := GenerateSessionCreationResponse(pl)
	assert.Nil(t, err)

	var response *responsePayload = new(responsePayload)
	marshalError := xml.Unmarshal([]byte(stringResponse), response)
	assert.Nil(t, marshalError)

	// Check SID
	assert.Equal(t, 40, len(response.SID))

	// Check Wait
	clientWait, _ := strconv.Atoi(pl.Wait)
	assert.True(t, response.Wait >= clientWait)

	// Check Requests
	clientHold, _ := strconv.Atoi(pl.Hold)
	assert.True(t, response.Requests > clientHold)
}

func TestValidGenerateSessionCreationsResposneReturnsPolling(t *testing.T) {
	pl := getSessionCreationPayload()
	pl.SID, _ = CreateNewSession()

	stringResponse, err := GenerateSessionCreationResponse(pl)
	assert.Nil(t, err)

	var response *responsePayload = new(responsePayload)
	xml.Unmarshal([]byte(stringResponse), response)

	assert.Equal(t, "15", response.Polling)
}
