package main

import (
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"net"
	"os"
)

func main() {
	hostname := os.Args[1]

	host, err := net.LookupIP(hostname)

	utils.ExitIfErrNotNil(err)

	for _, s := range host {
		fmt.Printf("%v\n", s)
	}
}
