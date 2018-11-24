package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/kodefish/go-presence-detection/api"
	"github.com/kodefish/go-presence-detection/detection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := api.NewRouter() // create routes

	allowedOrigins := handlers.AllowedOrigins([]string{"192.33.206.191"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"*"})

	// Launch server with CORS validation
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))

	for key, value := range detection.Scan() {
		fmt.Printf("%s - %s\n", key, value)
	}
}
