package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Database: Database{}}

// Route struct defining a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines list of routes of the API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddUser",
		"POST",
		"/register",
		controller.AddUser,
	},
	Route{
		"Auth",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	/*
		Route{
			"AddDevice",
			"POST",
			"/add/device",
			AuthenticationMiddleware(controller.AddDevice),
		},
		Route{
			"Devices",
			"GET",
			"/devices",
			AuthenticationMiddleware(controller.GetAllDevices),
		},
	*/
}

// NewRouter function configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Println(route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
