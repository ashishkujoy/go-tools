package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go greet(i)
	}

	waitGroup.Wait()
	c := make(chan int)
	go writeChannel(c, 10)
	defer close(c)
	time.Sleep(1 * time.Second)
	println("=================")
	_ = <-c
	time.Sleep(100 * time.Millisecond)
}

func greet(index int) {
	time.Sleep(time.Second)
	fmt.Println(fmt.Sprintf("Hello the index is %d", index))
	waitGroup.Done()
}

func writeChannel(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	fmt.Println(x)
	//close(c)
	fmt.Println(x)
}

