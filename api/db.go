package api

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
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

// GetUsers returns a list of all the users in the DB
func (db Database) GetUsers() Users {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	collection := session.DB(DBNAME).C(COLLECTION)
	results := Users{}

	if err = collection.Find(nil).All(&results); err != nil {
		log.Fatal(err)
	}

	fmt.Println(results)

	return results
}
