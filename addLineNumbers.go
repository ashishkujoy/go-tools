package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flags := os.Args

	if len(flags) == 0 {
		fmt.Printf("usage: addLineNumbers <files>\n")
		os.Exit(1)
	}

	for _, filename := range flags {
		fmt.Println("Processing:", filename)
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, line)
		}

		newContent := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
