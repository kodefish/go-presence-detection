package detection

// IP is a IP. duh.
type IP string

// MAC is a MAC. I can't believe I'm writing this shit.
type MAC string

// IPandMAC is, surprisingly, and IP and a MAC
type IPandMAC map[IP]MAC

// MACandAbscences maps a MAC address to the number of times it was absent from the call
type MACandAbscences map[MAC]int
