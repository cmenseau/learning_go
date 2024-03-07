package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var conv = func(b byte) string {
		return fmt.Sprint(int(b))
	}
	return conv(ip[0]) + "." + conv(ip[1]) + "." + conv(ip[2]) + "." + conv(ip[3])

	//return fmt.Sprint(int(ip[0])) + "." + fmt.Sprint(int(ip[1])) + "." + fmt.Sprint(int(ip[2])) + "." + fmt.Sprint(int(ip[3]))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
