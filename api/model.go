package api

import (
	"github.com/kodefish/go-presence-detection/detection"
)

type User struct {
	ID       string          `bson:"_id" json:"_id"`
	Name     string          `bson:"username" json:"username"`
	Password string          `bson:"password" json:"password"`
	Devices  []detection.MAC `bson:"devices" json:"devices"`
}

type UserDevices struct {
	Devices []detection.MAC `bson:"devices" json:"devices"`
}

// Users list of users
type Users []User

// JwtToken {token: "jwt token"}
type JwtToken struct {
	Token string `json:"token"`
}

// Exception server exception
type Exception struct {
	Message string `json:"message"`
}
