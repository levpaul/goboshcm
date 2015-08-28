package conn

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Body string `xml:"rid,attr"`
}

// The real workhorse
func postHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Println("Erorr reading body from request!")
	}
	defer r.Body.Close()

	var payload Payload
	xml.Unmarshal(body, payload)

	fmt.Println("P:", payload)

	/*
			   Overview
			    1. Parse the xml in the body (400 on bad parse) https://golang.org/pkg/encoding/xml/
			    2. Check for sid -> if exists
		          - get Session
		          if can't find, return 404 else:

			    3. Else create new session
		       - check RID (error if not exist)
		       - check to (error if not exist)
		       - check white/blacklist TODO: ADD LATER
		       - get wait (and set a default of 3)
		       - get xml:lang - default to 'en'
		       - get innactivity

	*/

}
