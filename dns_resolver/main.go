package main

import (
	"fmt"
	"net"
)

func main() {
	ips, _ := net.LookupIP("google.com")
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
