package conn

import "net/http"

const (
	ALLOWED_METHODS = "GET, POST, OPTIONS"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	getCommonHeaders(w, r)
	w.Header().Add("Access-Control-Allow-Methods", ALLOWED_METHODS)
	w.WriteHeader(200)
}
