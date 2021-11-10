package main

import (
	"fmt"
	sharedRPC "github.com/ashishkujoy/go-tools/sharedRpc"
	"github.com/ashishkujoy/go-tools/utils"
	"net/rpc"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		os.Exit(100)
	}

	connectionAddress := arguments[1]
	c, err := rpc.Dial("tcp", connectionAddress)
	utils.ExitIfErrNotNil(err)

	args := sharedRPC.MyInts{A1: 17, A2: 18, S1: true}
	var reply int

	err = c.Call("MyInterface.Add", args, &reply)
	utils.ExitIfErrNotNil(err)
	fmt.Printf("Reply (Add): %d\n", reply)

	err = c.Call("MyInterface.Subtract", args, &reply)
	utils.ExitIfErrNotNil(err)
	fmt.Printf("Reply (Subtract): %d\n", reply)
}
