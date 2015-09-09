package sessions

import (
	"encoding/xml"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/levilovelock/goboshcm/common"
	"github.com/nu7hatch/gouuid"
)

// TODO: Put this into config
const (
	MIN_WAIT     = 15
	MAX_WAIT     = 60
	MIN_REQUESTS = 2
	MAX_REQUESTS = 2
)

type Session struct {
	Authenticated bool
	Wait          int
	Requests      int
}

var (
	sessions map[string]*Session
)

func InitialiseSessionsPool() error {
	rand.Seed(int64(time.Now().Nanosecond()))
	sessions = make(map[string]*Session)
	return nil
}

func CreateNewSession() (string, error) {
	sid, err := generateNewSid()
	if err != nil {
		return "", err
	}
	sessions[sid] = new(Session)

	return sid, nil
}

func SessionExists(sid string) bool {
	return sessions[sid] != nil
}

func generateNewSid() (string, error) {
	// Simply appending two uuids together, stripping the dashes
	// and trimming to 40 chars
	sid := ""
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	sid += u4.String()
	u4, err = uuid.NewV4()
	if err != nil {
		return "", err
	}
	sid += u4.String()
	sid = strings.Replace(sid, "-", "", -1)

	return sid[0:40], nil
}

func GenerateSessionCreationResponse(p *common.Payload) (string, error) {
	returnPayload := new(common.Payload)

	// Set SID
	if p.SID == "" || sessions[p.SID] == nil {
		return "", errors.New("Error creating session response - no SID for session")
	}
	returnPayload.SID = p.SID
	clientSession := sessions[p.SID]

	// Set Wait
	clientWait, waitParseErr := strconv.Atoi(p.Wait)
	if waitParseErr != nil {
		return "", errors.New("Error parsing wait from client")
	}
	clientSession.Wait = common.Min(common.Max(clientWait, MIN_WAIT), MAX_WAIT)
	returnPayload.Wait = strconv.Itoa(clientSession.Wait)

	// Set Requests
	clientHold, holdParseErr := strconv.Atoi(p.Hold)
	if holdParseErr != nil {
		return "", errors.New("Error parsing hold from client")
	}
	if clientHold+1 > MAX_REQUESTS || clientHold+1 < MIN_REQUESTS {
		return "", errors.New("Unsupported hold value from client")
	}
	clientSession.Requests = common.Min(common.Max(clientHold, MIN_REQUESTS), MAX_REQUESTS)
	returnPayload.Requests = strconv.Itoa(clientSession.Requests)

	data, err := xml.Marshal(returnPayload)
	if err != nil {
		return "", errors.New("Error marshalling return payload")
	}
	return string(data), nil
}
