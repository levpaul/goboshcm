package sessions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
