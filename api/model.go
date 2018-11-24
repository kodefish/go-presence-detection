package api

// Device Type enum
type DeviceType int

const (
	Phone    DeviceType = 0
	Computer DeviceType = 1
	Tablet   DeviceType = 2
)

type User struct {
	ID      int     `bson:"_id"`
	Name    string  `json:"username"`
	Devices Devices `json:"devices"`
}
type Users []User

// TODO are these really strings or can we use smth else ?
type MAC string
type IP string
type Device struct {
	ID      int        `bson:"_id"`
	MACAddr MAC        `json:"mac_address"`
	IPAddr  IP         `json:"ip_address"`
	Name    string     `json:"device_name"`
	Type    DeviceType `json:"device_type"`
	Mobile  bool       `json:"mobile"`
}
type Devices []Device
