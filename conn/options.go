package conn

import "net/http"

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Length", "0")
	w.WriteHeader(200)
}
