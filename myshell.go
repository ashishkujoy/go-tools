package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")
		if len(words) > 0 {
			command := words[0]
			switch command {
			case "exit":
				fmt.Println("Bye Bye ........")
				os.Exit(0)
			case "version":
				fmt.Println("2.01")
			default:
				fmt.Println(text)
			}
			fmt.Print("> ")
		}
	}
}
