package detection

import (
	"log"
	"os/exec"
	"strings"
)

// Scan scans.
// That's it.
// Ta mère.
func Scan() (ipandmac IPandMAC) {
	ipandmac = make(IPandMAC)

	output, err := exec.Command("/bin/bash", "-c", "sudo arp-scan -l").Output()
	if err != nil {
		log.Fatal(err)
	}

	asStrings := strings.Split(string(output[:]), "\n")
	asStrings = asStrings[2 : len(asStrings)-4]
	for _, line := range asStrings {
		asArray := strings.Split(line, "\t")
		ipandmac[IP(asArray[0])] = MAC(strings.Split(asArray[1], " ")[0])
	}

	return
}
