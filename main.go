package main

import (
	"log"
	"net/http"

	"github.com/levilovelock/goboshcm/conn"
)

func main() {
	r := conn.ConstructRouter()
	log.Println("Starting GoBOSHCM server, listenining on :5280")
	http.ListenAndServe(":5280", r)
}
