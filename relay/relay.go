package relay

import (
	"net"
)

// IsICloudPrivateRelayAddress checks if the ip address is in the list of iCloud private relay ip networks
func IsICloudPrivateRelayAddress(ipAddress string) bool {
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

// ICloudPrivateRelay returns the csv row data from the egress ip ranges if the provided ip address is present
func ICloudPrivateRelay(ipAddress string) (egress, error) {
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
