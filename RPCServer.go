package main

import (
	sharedRPC "github.com/ashishkujoy/go-tools/sharedRpc"
	"github.com/ashishkujoy/go-tools/utils"
	"net"
	"net/rpc"
)

func main() {
	PORT := ":1234"
	myInterface := new(MyInterface)
	rpc.Register(myInterface)

	t, err := net.ResolveTCPAddr("tcp", PORT)
	utils.ExitIfErrNotNil(err)

	l, err := net.ListenTCP("tcp", t)
	defer l.Close()
	utils.ExitIfErrNotNil(err)

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(c)
	}
}

type MyInterface int

func (t *MyInterface) Add(arguments *sharedRPC.MyInts, reply *int) error {
	s1 := 1
	s2 := 1

	if arguments.S1 == true {
		s1 = -1
	}

	if arguments.S2 == true {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) + s2*int(arguments.A2)
	return nil
}

func (t *MyInterface) Subtract(arguments *sharedRPC.MyInts, reply *int) error {
	s1 := 1
	s2 := 1

	if arguments.S1 == true {
		s1 = -1
	}

	if arguments.S2 == true {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) - s2*int(arguments.A2)
	return nil
}