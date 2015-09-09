package conn

import (
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/levilovelock/goboshcm/common"
	"github.com/levilovelock/goboshcm/sessions"

	"encoding/xml"
)

// The real workhorse
func postHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)

	body, readErr := ioutil.ReadAll(r.Body)

	if readErr != nil {
		log.Infoln("Error reading body from request!")
	}
	defer r.Body.Close()

	var payload common.Payload

	err := xml.Unmarshal(body, &payload)
	if err != nil {
		log.Infoln("Failed to unmarshal xml payload:", err.Error())
		w.WriteHeader(400)
		return
	}

	if payload.SID == "" {
		// Here we are expecting a session creation request
		if validationErr := common.ValidatePayloadForSessionCreation(&payload); validationErr != nil {
			log.Infoln("Error validating payload for sessions creation:", validationErr.Error())
			w.WriteHeader(400)
			return
		}

		var sessionErr error
		payload.SID, sessionErr = sessions.CreateNewSession()
		if sessionErr != nil {
			log.Infoln("Error creating session,", err.Error())
			w.WriteHeader(500)
			return
		}

		// Here we need to populate the xml response and return it
		responseBody, sessionCreationErr := sessions.GenerateSessionCreationResponse(&payload)
		if sessionCreationErr != nil {
			log.Debug("Error creating session:", sessionCreationErr)
			w.WriteHeader(500)
			return
		}
		w.Write([]byte(responseBody))
		w.WriteHeader(200)
		return

	} else if !sessions.SessionExists(payload.SID) {
		log.Infoln("Invalid sid supplied for request")
		w.WriteHeader(400)
		return
	}
}
