package main

import (
	"fmt"
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

	router := api.NewRouter() // create routes
	// Launch server with CORS validation
	log.Fatal(http.ListenAndServe(":"+port, router))

	for key, value := range detection.Scan() {
		fmt.Printf("%s - %s\n", key, value)
	}
}
