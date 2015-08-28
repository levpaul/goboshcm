package conn

import "net/http"

const (
	ALLOWED_METHODS = "GET, POST, OPTIONS"
)

var (
	ALLOWED_ORIGINS string
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	if origin != "" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}

	w.Header().Add("Access-Control-Allow-Methods", ALLOWED_METHODS)
	w.WriteHeader(200)
}
