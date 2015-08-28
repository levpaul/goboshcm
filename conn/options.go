package conn

import "net/http"

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
