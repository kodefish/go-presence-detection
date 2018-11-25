package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kodefish/go-presence-detection/crypto"

	"github.com/gorilla/context"
)

const (
	secretSigningKey = "superS3cretKeyToSignJWT"
)

type Controller struct {
	Database Database
}

// AuthenticationMiddleware makes sure the request has a valid JWT token
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Auth")
		if authorizationHeader := r.Header.Get("authorization"); authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(secretSigningKey), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					log.Println("TOKEN WAS VALID")
					context.Set(r, "decoded", token.Claims)
					next(w, r)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}

		}
	})

}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Database.GetUsers() // list of all products
	// log.Println(products)
	data, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "192.33.206.191")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func sendJwtToken(user User, w http.ResponseWriter) {
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})

	// Send token response
	if tokenString, err := token.SignedString([]byte(secretSigningKey)); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/jwt; charset=UTF-8")
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	}

}

func parseUserFromFrom(f *http.Request) User {
	dumpRequest(f)
	if err := f.ParseForm(); err != nil {
		return User{}
	}

	var user User
	for key, value := range f.Form {
		switch key {
		case "username":
			user.Name = value[0]
		case "password":
			user.Password = value[0]
		}
	}
	log.Println("received user", user)
	return user
}

// GetToken Autenthifies user
func (c *Controller) GetToken(w http.ResponseWriter, r *http.Request) {
	// Get user
	dumpRequest(r)

	user := parseUserFromFrom(r)

	// Get the user's password from the db
	var dbUser User
	c.Database.GetUserByUsername(user.Name, &dbUser)

	// Check that the password matches
	if crypto.ComparePasswords(dbUser.Password, user.Password) {
		sendJwtToken(dbUser, w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// AddUser adds a user if the username is not already taken
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	newUsr := parseUserFromFrom(r)

	// Make sure username does not already exist
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if c.Database.GetUserByUsername(newUsr.Name, &newUsr) {
		log.Println("User exists!")

		// If sucessfull -> user exists -> reply with user already exists
		w.WriteHeader(http.StatusTeapot)
		json.NewEncoder(w).Encode(Exception{Message: "Username is already assigned"})
	} else {
		log.Println("User does not exist!")
		if usr, err := c.Database.AddUser(newUsr); err {
			log.Println("User added!")
			sendJwtToken(usr, w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (c *Controller) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) WhosHome(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	macs []detection.MAC := detection.GetMacs()
	usersHome := make(map[string]string) // username -> one of their device's MAC

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, mac := range macs {
		if var dbUser User; c.Database.getUserByDevice(mac, &dbUser) {
			usersHome[dbUser] = mac
		}
	}

	json.NewEncoder(w).Encode(usersHome)
}

func dumpRequest(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error dumping request:", err)
		return
	}
	log.Println(string(output))
}
