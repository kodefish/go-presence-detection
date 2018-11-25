package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kodefish/go-presence-detection/api"
	"github.com/kodefish/go-presence-detection/detection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Launch arp scanner
	go detection.Update()

	// Init REST server
	router := api.NewRouter() // create routes
	// Launch server with CORS validation
	log.Println("Starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
