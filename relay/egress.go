package relay

import (
	"encoding/binary"
	"encoding/csv"
	"net"
	"net/http"
)

type location struct {
	CountryCode string
	State       string
	City        string
}

var locations = map[string]location{}
var ipAddresses = map[string]string{}

func mapAddresses(data [][]string) (ipAddresses map[string]string, locations map[string]location, err error) {
	locations = map[string]location{}
	ipAddresses = map[string]string{}
	for _, row := range data {
		if len(row) < 4 {
			continue
		}
		cidr := row[0]
		locations[cidr] = location{
			CountryCode: row[1],
			State:       row[2],
			City:        row[3],
		}

		expanded, err := expandCIDR(cidr)
		if err != nil {
			return ipAddresses, locations, err
		}
		for _, ipAddress := range expanded {
			ipAddresses[ipAddress] = cidr
		}
	}
	return ipAddresses, locations, nil
}

func addresses() (data [][]string, err error) {
	data, err = readCSVFromUrl("https://mask-api.icloud.com/egress-ip-ranges.csv")
	return
}

func readCSVFromUrl(url string) (data [][]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err = reader.ReadAll()
	return
}

func expandCIDR(CIDR string) ([]string, error) {
	// convert string to IPNet struct
	_, ipv4Net, err := net.ParseCIDR(CIDR)
	if err != nil {
		return nil, err
	}

	// convert IPNet struct mask and address to uint32
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	ips := []string{}
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips = append(ips, ip.String())
	}
	return ips, nil
}
