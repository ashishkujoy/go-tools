package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Epoch time: %v\n", now.Unix())
	fmt.Printf("Formated time %v\n", now.Format(time.RFC3339))
	time.Sleep(time.Second)

	fmt.Printf("Time difference: %v\n", time.Now().Sub(now))

	file, _ := os.Open("/dev/random")
	defer file.Close()
	var seed int64
	binary.Read(file, binary.LittleEndian, &seed)

	fmt.Printf("The seed : %v\n", seed)
	rand.Seed(seed)
	fmt.Printf("The random number %v\n",rand.Int())
}
