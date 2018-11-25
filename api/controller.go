package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kodefish/go-presence-detection/crypto"
	"github.com/kodefish/go-presence-detection/detection"

	"github.com/gorilla/context"
)

const (
	secretSigningKey = "superS3cretKeyToSignJWT"
)

type Controller struct {
	Database Database
}

func dumpRequest(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error dumping request:", err)
		return
	}
	log.Println(string(output))
}

func getUserIDFromContext(r *http.Request) string {
	tokenClaims := context.Get(r, "tokenClaims")
	return tokenClaims.(jwt.MapClaims)["user_id"].(string)
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
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					context.Set(r, "tokenClaims", claims)
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

func parseUserFromJSON(w http.ResponseWriter, r *http.Request) User {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// Unmarshal
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return user
}

// GetToken Autenthifies user
func (c *Controller) GetToken(w http.ResponseWriter, r *http.Request) {

	// Get user
	user := parseUserFromJSON(w, r)

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
	newUsr := parseUserFromJSON(w, r)

	// Make sure username does not already exist
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if c.Database.GetUserByUsername(newUsr.Name, &newUsr) {
		// If sucessfull -> user exists -> reply with user already exists
		w.WriteHeader(http.StatusTeapot)
		json.NewEncoder(w).Encode(Exception{Message: "Username is already assigned"})
	} else {
		if usr, err := c.Database.AddUser(newUsr); err {
			sendJwtToken(usr, w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetAllDevices returns the list of devices a user owns
func (c *Controller) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	var user User
	if c.Database.GetUserByID(getUserIDFromContext(r), &user) {
		//w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserDevices{user.Devices})
		//json.NewEncoder(w).Encode(user.Devices)
	}
}

func (c *Controller) AddDevice(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	userID := getUserIDFromContext(r)
	log.Println(r.RemoteAddr, userID)
	userIP := detection.IP((strings.Split(r.RemoteAddr, ":")[0]))
	userMAC := detection.GetMACfromIP(userIP)
	// Check that MAC is not already in DB
	if userMAC != "" {
		if c.Database.MACInDB(userMAC) {
			json.NewEncoder(w).Encode(Exception{Message: "This device is already in the DB"})
			return
		}
		var user User
		if c.Database.GetUserByID(userID, &user) {
			user.Devices = append(user.Devices, userMAC)
			if c.Database.UpdateUserById(userID, user) {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				json.NewEncoder(w).Encode(UserDevices{user.Devices})
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusTeapot)
}

func setContains(set []string, elem string) bool {
	for _, item := range set {
		if item == elem {
			return true
		}
	}
	return false
}

func (c *Controller) WhosHome(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	macs := detection.GetMACs()
	usersHome := make(map[string]detection.MAC) // username -> one of their device's MAC

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var connectedUser []string
	for _, mac := range macs {
		var dbUser User
		if c.Database.getUserByDevice(mac, &dbUser) {
			if !setContains(connectedUser, dbUser.name) {
				connectedUser = append(connectedUser, dbUser.Name)
				usersHome[dbUser.Name] = mac
			}
		}
	}

	json.NewEncoder(w).Encode(connectedUser)
}

func getNewDevices(set []detection.MAC, elem detection.MAC) (newDevices []detection.MAC) {
	newDevices = make([]detection.MAC)
	for _, item := range set {
		if item != elem {
			newDevices.push(item)
		}
	}
	return
}

func (c *Controller) DeleteDevices(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	userID := getUserIDFromContext(r)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var dev struct {
		Devices []detection.MAC `json:"devices" bson:"devices"`
	}
	err = json.Unmarshal(b, &dev)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user User
	if c.Database.GetUserByID(userID, &user) {
		user.Devices = getNewDevices(user.Devices, dev.Devices[0])
		log.Println(user.Devices)
		if c.Database.UpdateUserById(userID, user) {
			json.NewEncoder(w).Encode(user.Devices)
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)

}
