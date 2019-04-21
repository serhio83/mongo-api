package main

import (
	"log"
	"net/http"
	"os"

	"github.com/serhio83/mongo-api/pkg/handlers"
	"github.com/serhio83/mongo-api/pkg/version"
)

func main() {
	log.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

