package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	blockSize := flag.Int("bs", 0, "Block Size")
	count := flag.Int("count", 0, "Counter")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(-1)
	}

	if *blockSize < 0 || *count < 0 {
		fmt.Println("Count or/and Byte Size < 0!")
		os.Exit(-1)
	}

	filename := flags[0]

	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}
	destination, err := os.Create(filename)
	defer destination.Close()
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	buf := make([]byte, *blockSize)
	for i := 0; i < *count; i++ {
		createBytes(&buf, *blockSize)
		if _, err := destination.Write(buf); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		buf = nil
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + max
}

func createBytes(buf *[]byte, count int) {
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		b := byte(random(0, 9))
		*buf = append(*buf, b)
	}
}
