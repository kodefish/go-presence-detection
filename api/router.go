package api

import (
	"time"

	"github.com/husobee/vestigo"
)

var controller = &Controller{Database: Database{}}

// NewRouter function configures a new router to the API
func NewRouter() *vestigo.Router {
	router := vestigo.NewRouter()
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "test.com"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Header", "X-Y-Header"},
		MaxAge:           3600 * time.Second,
		AllowHeaders:     []string{"X-Header", "X-Y-Header", "Authorization"},
	})

	// Define Routes
	router.Get("/", controller.Index)
	router.Post("/register", controller.AddUser)
	router.Post("/get-token", controller.GetToken)
	router.Get("/devices", controller.GetAllDevices, AuthenticationMiddleware)
	router.Get("/connected-users", controller.WhosHome, AuthenticationMiddleware)
	router.Get("/add-device", controller.AddDevice, AuthenticationMiddleware)
	router.Post("/delete-device", controller.DeleteDevices, AuthenticationMiddleware)
	return router
}
