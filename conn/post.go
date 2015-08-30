package conn

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/levilovelock/goboshcm/sessions"

	"encoding/xml"
)

// The real workhorse
func postHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)

	body, readErr := ioutil.ReadAll(r.Body)

	if readErr != nil {
		log.Println("Error reading body from request!")
	}
	defer r.Body.Close()

	var payload Payload

	err := xml.Unmarshal(body, &payload)
	if err != nil {
		log.Println("Failed to unmarshal xml payload:", err.Error())
		w.WriteHeader(400)
		return
	}

	if payload.SID == "" {

		if validationErr := validatePayloadForSessionCreation(&payload); validationErr != nil {
			log.Println("Error validating payload for sessions creation:", validationErr.Error())
			w.WriteHeader(400)
			return
		}

		var sessionErr error
		payload.SID, sessionErr = sessions.CreateNewSession()
		if sessionErr != nil {
			log.Println("Error creating session,", err.Error())
			w.WriteHeader(500)
			return
		}

	} else if !sessions.SessionExists(payload.SID) {
		log.Println("Invalid sid supplied for request")
		w.WriteHeader(400)
		return
	}

	/*
			   Overview
			    1. Parse the xml in the body (400 on bad parse) https://golang.org/pkg/encoding/xml/
			    2. Check for sid -> if exists
		          - get Session
		          if can't find, return 400 else:

			    3. Else create new session
		       - check RID (error if not exist)
		       - check to (error if not exist)
		       - check white/blacklist TODO: ADD LATER
		       - get wait (and set a default of 3)
		       - get xml:lang - default to 'en'
		       - get innactivity

	*/
}
