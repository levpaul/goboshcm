package conn

import "net/http"

const (
	ALLOWED_METHODS = "GET, POST, OPTIONS"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Methods", ALLOWED_METHODS)
	w.WriteHeader(200)
}
