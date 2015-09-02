package common

import (
	"encoding/xml"
	"errors"
)

type Payload struct {
	XMLName     xml.Name `xml:"body"`
	RID         string   `xml:"rid,attr"`
	To          string   `xml:"to,attr"`
	XMLNS       string   `xml:"xmlns,attr"`
	XMLLang     string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Wait        string   `xml:"wait,attr"`
	Hold        string   `xml:"hold,attr"`
	Content     string   `xml:"content,attr"`
	Version     string   `xml:"ver,attr"`
	XMPPVersion string   `xml:"urn:xmpp:xbosh version,attr"`
	XMLNSXMPP   string   `xml:"xmlns xmpp,attr"`
	Route       string   `xml:"route,attr"`
	SID         string   `xml:"sid,attr"`
}

func ValidatePayloadForSessionCreation(pl *Payload) error {
	// Check that mandatory fields are set
	// TODO: Added whitelist/blacklist validations
	if pl.To == "" || pl.XMLLang == "" || pl.Version == "" || pl.Wait == "" || pl.Hold == "" || pl.RID == "" {
		return errors.New("Invalid xml body for session creation")
	} else {
		return nil
	}
}
