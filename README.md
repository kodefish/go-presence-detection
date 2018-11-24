# Sensorless Presence Detection System

## Dependencies
### Library to handle jwt authentication 
$ go get "github.com/dgrijalva/jwt-go"

### mgo library for handling MongoDB
$ go get "gopkg.in/mgo.v2"

### MongoDB
The presence detection system uses MongoDB as a database.

## Repo structures
* api/ - contains all api related stuff
    + controller.go - Defines handler methods for endpoints
    + model.go      - User models
    + db.go         - Methods to interact with db
    + router.go     - Defines routes and endpoints
* README.md         - Readme for documentation
* main.go           - Entry point for api
