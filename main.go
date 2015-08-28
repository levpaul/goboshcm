package main

import (
	"log"
	"net/http"

	"github.com/levilovelock/goboshcm/conn"
)

func main() {
	r := conn.ConstructRouter()
	log.Println("Starting GoBOSHCM server, listenining on :5280")
	err := http.ListenAndServe(":5280", r)
	if err != nil {
		log.Fatal("There was an error starting the server:", err.Error())
	}
}
