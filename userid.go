package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) == 1 {
		currentUserId := os.Getuid()
		fmt.Printf("Current UserId: %d\n", currentUserId)
		os.Exit(0)
	}

	username := os.Args[1]
	lookup, err := user.Lookup(username)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("UserId of %s : %s\n", username, lookup.Uid)
}
