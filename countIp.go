package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	columnNumber := flag.Int("col", 1, "Column")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if *columnNumber < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	myIPs := make(map[string]int)

	for _, filename := range flags {
		fmt.Println("\t\t", filename)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				continue
			}

			data := strings.Fields(line)
			ip := data[*columnNumber-1]
			_, ok := myIPs[ip]
			if ok {
				myIPs[ip] = myIPs[ip] + 1
			} else {
				myIPs[ip] = 1
			}
		}
	}
	for key, _ := range myIPs {
		fmt.Printf("%s %d\n", key, myIPs[key])
	}
}
