package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var result string

	for _, value := range ip {
		result += fmt.Sprintf("%v.", value)
	}

	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v IP: %v\n", name, ip)
	}
}
