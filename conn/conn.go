package conn

import "github.com/gorilla/mux"

func ConstructRouter() *mux.Router {
	r := mux.NewRouter()

	r.Methods("OPTIONS").Path("/http-bind").HandlerFunc(optionsHandler)
	r.Methods("GET").Path("/http-bind").HandlerFunc(getHandler)
	r.Methods("POST").Path("/http-bind").HandlerFunc(postHandler)

	return r
}
