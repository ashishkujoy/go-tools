package main

import (
	"github.com/ashishkujoy/go-tools/utils"
	"net"
	"os"
)

func main() {
	ip := os.Args[1]
	parseIP := net.ParseIP(ip)

	if parseIP == nil {
		println("Invalid ip address....")
		os.Exit(1)
	}

	addr, err := net.LookupAddr(ip)

	utils.ExitIfErrNotNil(err)

	for _, address := range addr {
		println(address)
	}
}
