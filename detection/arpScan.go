package detection

import (
	"log"
	"os/exec"
	"strings"
)

// Scan scans.
// That's it.
// Ta mère.
func Scan() (IPandMAC map[string]string) {
	IPandMAC = make(map[string]string)

	output, err := exec.Command("/bin/bash", "-c", "arp-scan -l").Output()
	if err != nil {
		log.Fatal(err)
	}

	asStrings := strings.Split(string(output[:]), "\n")
	asStrings = asStrings[2 : len(asStrings)-4]
	for _, line := range asStrings {
		asArray := strings.Split(line, "\t")
		IPandMAC[asArray[0]] = asArray[1]
	}

	return
}
