package detection

import (
	"time"
)

const (
	maxConsecutiveAbsences = 3
	sleepTime              = 10
)

var (
	currentlyConnected = make(IPandMAC)
	absents            = make(MACandAbscences)
	responders         = make(IPandMAC)
)

// Update is an infinite loop that constantly updates the list of users connected to the network
func Update() {
	for {
		takeAttendance()
		time.Sleep(sleepTime)
	}
}

func takeAttendance() {
	// 1. Taking new presences
	responders = Scan()

	// 2. Updating new presences
	for ip, mac := range responders {
		if _, isContained := currentlyConnected[ip]; !isContained {
			// ip is coming for the first time
			currentlyConnected[ip] = mac
			// it has now been absent for 0 turns
			absents[mac] = 0
		}
	}

	// 3. Updating abscences
	for ip, mac := range currentlyConnected {
		if _, isComingBack := responders[ip]; !isComingBack {
			// ip is not coming back this time
			if consecutiveAbsences, _ := absents[mac]; consecutiveAbsences < maxConsecutiveAbsences {
				// Adding one shot
				absents[mac]++
			} else {
				// Expelling that user
				delete(currentlyConnected, ip)
				delete(absents, mac)
			}
		}
	}
}
