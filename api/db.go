package api

import (
	"log"

	"github.com/kodefish/go-presence-detection/crypto"
	"github.com/satori/go.uuid"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Database object
type Database struct{}

const (
	// SERVER Mongodb server
	SERVER = "mongodb://127.0.0.1:27017"
	// DBNAME DB Name
	DBNAME = "spds"
	// COLLECTION Name
	COLLECTION = "users"
)

func getSessionAndCollection() (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	collection := session.DB(DBNAME).C(COLLECTION)
	return session, collection
}

// GetUsers returns a list of all the users in the DB
func (db Database) GetUsers() Users {
	session, collection := getSessionAndCollection()
	defer session.Close()
	results := Users{}

	if err := collection.Find(nil).All(&results); err != nil {
		log.Fatal(err)
	}

	return results
}

// GetUserByID retrieves the user with the corresponding id from the db
func (db Database) GetUserByID(id int, result *User) bool {
	session, collection := getSessionAndCollection()
	defer session.Close()

	err := collection.Find(bson.M{"_id": id}).One(result)
	return err == nil
}

// GetUserByUsername retrieves the user with the corresponding username from the db into the user, return true if success
func (db Database) GetUserByUsername(username string, result *User) bool {
	session, collection := getSessionAndCollection()
	defer session.Close()

	err := collection.Find(bson.M{"username": username}).One(result)
	return err == nil
}

// AddUser adds a user to the db
func (db Database) AddUser(user User) (User, bool) {
	session, collection := getSessionAndCollection()
	defer session.Close()
	// Generate new user id
	id, err := uuid.NewV4()
	// Generate salted password
	saltedPassword := crypto.HashAndSalt(user.Password)

	// Store user in db - return error. But should go smoothly, right?
	usr := User{
		ID:       id.String(),
		Name:     user.Name,
		Password: saltedPassword,
	}
	err = collection.Insert(usr)
	return usr, err == nil
}
