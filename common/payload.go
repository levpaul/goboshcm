package common

import (
	"encoding/xml"
	"errors"
)

const (
	DEF_XMLNS = "http://jabber.org/protocol/httpbind"
)

type Payload struct {
	XMLName     xml.Name `xml:"body"`
	RID         string   `xml:"rid,attr,omitempty"`
	To          string   `xml:"to,attr,omitempty"`
	XMLNS       string   `xml:"xmlns,attr,omitempty"`
	XMLLang     string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Wait        string   `xml:"wait,attr,omitempty"`
	Hold        string   `xml:"hold,attr,omitempty"`
	Content     string   `xml:"content,attr,omitempty"`
	Version     string   `xml:"ver,attr,omitempty"`
	XMPPVersion string   `xml:"urn:xmpp:xbosh version,attr,omitempty"`
	XMLNSXMPP   string   `xml:"xmlns xmpp,attr,omitempty"`
	Route       string   `xml:"route,attr,omitempty"`
	SID         string   `xml:"sid,attr,omitempty"`
	Requests    string   `xml:"requests,attr,omitempty"`
	Polling     string   `xml:"polling,attr,omitempty"`
}

func NewPayload() *Payload {
	return &Payload{XMLNS: DEF_XMLNS}
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
