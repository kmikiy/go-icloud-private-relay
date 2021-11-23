package relay

import (
	"encoding/csv"
	"net"
	"net/http"
)

type egress struct {
	IPNet       net.IPNet
	CIDR        string
	CountryCode string
	State       string
	City        string
}

var ipNets = []egress{}

func mapAddresses(data [][]string) (ipNets []egress, err error) {
	for _, row := range data {
		_, ipnet, err := net.ParseCIDR(row[0])
		if err != nil {
			return ipNets, err
		}
		if ipnet == nil {
			continue
		}
		if len(row) < 4 {
			continue
		}
		ipNets = append(ipNets, egress{
			IPNet:       *ipnet,
			CIDR:        row[0],
			CountryCode: row[1],
			State:       row[2],
			City:        row[3],
		})
	}
	return ipNets, nil
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
