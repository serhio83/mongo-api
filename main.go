package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/serhio83/mongo-api/pkg/handlers"
	"github.com/serhio83/mongo-api/pkg/version"
)

func main() {
	fmt.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s\n",
		version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()
	log.Printf("The service is listen on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
