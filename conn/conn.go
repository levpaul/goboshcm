package conn

import "github.com/gorilla/mux"

func ConstructRouter() *mux.Router {
	r := mux.NewRouter()

	r.Methods("OPTIONS").Path("/http-bind").HandlerFunc(optionsHandler)

	return r
}
