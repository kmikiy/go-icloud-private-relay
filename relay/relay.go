package relay

// IsICloudPrivateRelayAddress checks if the ip address is in the list of iCloud private relay ip networks
func IsICloudPrivateRelayAddress(ipAddress string) bool {
	loadRelayData()
	if _, ok := ipAddresses[ipAddress]; ok {
		return true
	}
	return false
}

// ICloudPrivateRelay returns the csv row data from the egress ip ranges if the provided ip address is present
func ICloudPrivateRelay(ipAddress string) (Location, error) {
	loadRelayData()
	if cidr, ok := ipAddresses[ipAddress]; ok {
		if location, ok := locations[cidr]; ok {
			return location, nil
		}
	}
	return Location{}, ErrNotFound
}

// loadRelayData loads relay informations from Apple.
func loadRelayData() {
	if len(locations) == 0 || len(ipAddresses) == 0 {
		a, _ := addresses()
		ipAddresses, locations, _ = mapAddresses(a)
	}
}
