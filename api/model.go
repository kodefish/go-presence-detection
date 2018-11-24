package api

// Device Type enum
type DeviceType int

const (
	Phone    DeviceType = 0
	Computer DeviceType = 1
	Tablet   DeviceType = 2
)

type User struct {
	ID int `bson:"_id" json:"_id"`

	Name    string  `bson:"username" json:"username"`
	Devices Devices `bson:"devices" json:"devices"`
}
type Users []User

// TODO are these really strings or can we use smth else ?
type MAC string
type IP string
type Device struct {
	ID      int        `bson:"_id" json:"_id"`
	MACAddr MAC        `bson:"mac_address" json:"mac_address"`
	IPAddr  IP         `bson:"ip_address" json:"ip_address"`
	Name    string     `bson:"device_name" json:"device_name"`
	Type    DeviceType `bson:"device_type" json:"device_type"`
	Mobile  bool       `bson:"mobile" json:"mobile"`
}
type Devices []Device
