package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kodefish/go-presence-detection/crypto"

	"github.com/dgrijalva/jwt-go"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// GetToken Autenthifies user
func (c *Controller) GetToken(w http.ResponseWriter, r *http.Request) {
	// Get user
	dumpRequest(r)
	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("received user", user)

	// Get the user's password from the db
	var dbUser User
	c.Database.GetUserByUsername(user.Name, &dbUser)

	// Check that the password matches
	if crypto.ComparePasswords(dbUser.Password, user.Password) {
		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": dbUser.ID,
		})

		// Send token response
		if tokenString, err := token.SignedString([]byte(secretSigningKey)); err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/jwt; charset=UTF-8")
			json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// AddUser adds a user if the username is not already taken
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	var newUsr User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUsr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Make sure username does not already exist
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if c.Database.GetUserByUsername(newUsr.Name, &newUsr) {
		log.Println("User exists!")

		// If sucessfull -> user exists -> reply with user already exists
		w.WriteHeader(http.StatusTeapot)
		json.NewEncoder(w).Encode(Exception{Message: "Username is already assigned"})
	} else {
		log.Println("User does not exist!")
		if c.Database.AddUser(newUsr) {
			log.Println("User added!")
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func dumpRequest(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error dumping request:", err)
		return
	}
	log.Println(string(output))
}
