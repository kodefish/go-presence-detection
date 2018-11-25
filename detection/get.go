package detection

import (
	"log"
)

// GetMACfromIP as surprising as it might sound, gets a MAC from an IP
func GetMACfromIP(ip IP) (mac MAC) {
	mac, _ = currentlyConnected[ip]
	log.Println(mac)
	return
}

// GetMACs well... take a guess
func GetMACs() (macs []MAC) {
	macs = make([]MAC, len(currentlyConnected))
	i := 0
	for _, mac := range currentlyConnected {
		macs[i] = mac
		i++
	}
	return
}
