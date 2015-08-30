package conn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationOfPayloadForSessionCreationSuccess(t *testing.T) {
	pl := getValidPayloadForSessionCreation()

	err := validatePayloadForSessionCreation(pl)
	assert.Nil(t, err)
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
