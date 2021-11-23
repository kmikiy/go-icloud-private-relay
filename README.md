# go-apple-private-relay

[![Go Reference](https://pkg.go.dev/badge/github.com/kmikiy/go-apple-private-relay/relay.svg)](https://pkg.go.dev/github.com/kmikiy/go-apple-private-relay/relay)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmikiy/go-apple-private-relay)](https://goreportcard.com/report/github.com/kmikiy/go-apple-private-relay)


Go package that detects iCloud Private Relay IP address based on Apple's [egress ip range](https://mask-api.icloud.com/egress-ip-ranges.csv) list. More information about how to prepare your network for iCloud Private Relay can be found [here](https://developer.apple.com/support/prepare-your-network-for-icloud-private-relay/).

## Installation

```
go get github.com/kmikiy/go-apple-private-relay
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/kmikiy/go-apple-private-relay/relay"
)

func main() {
	isApplePrivateRelayAddress := relay.IsApplePrivateRelayAddress("172.225.18.12")
	fmt.Println(isApplePrivateRelayAddress) // true

	isApplePrivateRelayAddress = relay.IsApplePrivateRelayAddress("142.251.39.14")
	fmt.Println(isApplePrivateRelayAddress) // false
}
```
