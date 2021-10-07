package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	message := os.Args[2]

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer file.Close()

	file.WriteString(message)

	fmt.Printf("File: %v, Message: %v\n", fileName, message)
}
