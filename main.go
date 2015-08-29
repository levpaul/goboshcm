package main

import (
	"log"
	"net/http"

	"github.com/levilovelock/goboshcm/conn"
	"github.com/levilovelock/goboshcm/sessions"
)

func main() {
	r := conn.ConstructRouter()
	sessionsErr := sessions.InitialiseSessionsPool()
	if sessionsErr != nil {
		log.Fatal("An error occurred starting the sessions pool,", sessionsErr.Error())
	}

	log.Println("Starting GoBOSHCM server, listenining on :5280")
	err := http.ListenAndServe(":5280", r)
	if err != nil {
		log.Fatal("There was an error starting the server:", err.Error())
	}
}
