package conn

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"encoding/xml"
)

type Payload struct {
	XMLName   xml.Name `xml:"body"`
	Rid       int      `xml:"rid,attr"`
	To        string   `xml:"to,attr"`
	Xmlns     string   `xml:"xmlns,attr"`
	XmlLang   string   `xml:"xmllang,attr"`
	Wait      int      `xml:"wait,attr"`
	Hold      int      `xml:"hold,attr"`
	Content   string   `xml:"content,attr"`
	Ver       string   `xml:"ver,attr"`
	XmppVer   string   `xml:"xmppversion,attr"`
	XmlnsXmpp string   `xml:"xmlnsxmpp,attr"`
	Route     string   `xml:"route,attr"`
}

// The real workhorse
func postHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)

	body, readErr := ioutil.ReadAll(r.Body)

	/* TODO: This is a hack because the encoding/xml library can't
	*  handle colons in the attr names. See the following isse
	*  https://github.com/golang/go/issues/11735 */
	body = sanitizeXml(body)

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

// TODO: Remove this function or optimise it
func sanitizeXml(b []byte) []byte {
	s := string(b)
	s = strings.Replace(s, "xml:lang=", "xmllang=", 1)
	s = strings.Replace(s, "xmpp:version=", "xmllang=", 1)
	s = strings.Replace(s, "xmlns:xmpp=", "xmlnsxmpp=", 1)
	return []byte(s)
}
