package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationOfPayloadForSessionCreationSuccess(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	err := ValidatePayloadForSessionCreation(pl)
	assert.Nil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnRID(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.RID = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnTo(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.To = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnVersion(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.Version = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnXMLLang(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.XMLLang = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnWait(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.Wait = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestPayloadValidationForSessionCreationFailsOnHold(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	pl.Hold = ""

	err := ValidatePayloadForSessionCreation(pl)
	assert.NotNil(t, err)
}

func TestNewPayloadHasCorrectXMLNS(t *testing.T) {
	pl := NewPayload()
	assert.Equal(t, pl.XMLNS, "http://jabber.org/protocol/httpbind")
}

func getValidPayloadForSessionCreation() *Payload {
	return &Payload{
		RID:     "2902921866",
		To:      "chat.mysite.com",
		Version: "1.6",
		XMLLang: "en",
		Wait:    "60",
		Hold:    "1",
	}
}
