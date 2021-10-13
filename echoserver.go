package main

import (
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a socket file.")
		os.Exit(100)
	}
	socketFile := arguments[1]
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		go echoServer(fd)
	}
}

func echoServer(c net.Conn) {
	for {
		buff := make([]byte, 1024)
		bytesRead, err := c.Read(buff)

		utils.ExitIfErrNotNil(err)

		data := buff[0:bytesRead]
		fmt.Printf("->: %v", string(data))
		_, err = c.Write(data)

		utils.ExitIfErrNotNil(err)
	}
}
