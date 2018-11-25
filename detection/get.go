package detection

func getMACs() (macs []MAC) {
	macs = make([]MAC, len(currentlyConnected))
	i := 0
	for _, mac := range currentlyConnected {
		macs[i] = mac
		i++
	}
	return
}

func getMACfromIP(ip IP) (mac MAC) {
	mac, _ = currentlyConnected[ip]
	return
}
