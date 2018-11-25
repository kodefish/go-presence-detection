package api

import "github.com/kodefish/go-presence-detection/detection"

// Device Type enum
type DeviceType int

const (
	Phone    DeviceType = 0
	Computer DeviceType = 1
	Tablet   DeviceType = 2
)

type User struct {
	ID       string  `bson:"_id" json:"_id"`
	Name     string  `bson:"username" json:"username"`
	Password string  `bson:"password" json:"password"`
	Devices  Devices `bson:"devices" json:"devices"`
}

// Users list of users
type Users []User

type Device struct {
	ID      int           `bson:"_id" json:"_id"`
	MACAddr detection.MAC `bson:"mac_address" json:"mac_address"`
	IPAddr  detection.IP  `bson:"ip_address" json:"ip_address"`
	Name    string        `bson:"device_name" json:"device_name"`
	Type    DeviceType    `bson:"device_type" json:"device_type"`
	Mobile  bool          `bson:"mobile" json:"mobile"`
}

// Devices list of devices
type Devices []Device

// JwtToken {token: "jwt token"}
type JwtToken struct {
	Token string `json:"token"`
}

// Exception server exception
type Exception struct {
	Message string `json:"message"`
}
