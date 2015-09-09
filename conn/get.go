package conn

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Just return a link to the XEP-124 spec
func getHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)
	_, err := w.Write([]byte(`<html>
  <body>
    <a href='http://www.xmpp.org/extensions/xep-0124.html'>XEP-0124</a> - BOSH
  </body>
</html>`))
	if err != nil {
		log.Infoln("Error writing response to GET request:", err.Error())
	}
}
