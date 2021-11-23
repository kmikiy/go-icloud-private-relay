package relay

import (
	"net"
)

// IsApplePrivateRelayAddress checks if the ip address is in the list of apple private relay ip networks
func IsApplePrivateRelayAddress(ipAddress string) bool {
	if len(ipNets) == 0 {
		a, _ := addresses()
		ipNets, _ = mapAddresses(a)
	}

	for _, egress := range ipNets {
		ip := net.ParseIP(ipAddress)
		if egress.IPNet.Contains(ip) {
			return true
		}
	}
	return false
}

// ApplePrivateRelay returns the csv row data from the egress ip ranges if the provided ip address is present
func ApplePrivateRelay(ipAddress string) (egress, error) {
	if len(ipNets) == 0 {
		a, _ := addresses()
		ipNets, _ = mapAddresses(a)
	}

	for _, egress := range ipNets {
		ip := net.ParseIP(ipAddress)
		if egress.IPNet.Contains(ip) {
			return egress, nil
		}
	}
	return egress{}, ErrNotFound
}
