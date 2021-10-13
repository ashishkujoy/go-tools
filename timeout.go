package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	duration := 2 * time.Second
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&waitGroup, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	waitGroup.Done()

	if timeout(&waitGroup, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

}

func timeout(waitGroup *sync.WaitGroup, t time.Duration) bool {
	tempChannel := make(chan int)

	go func() {
		defer close(tempChannel)
		waitGroup.Wait()
	}()

	select {
	case <-tempChannel:
		{
			println("all go routine done...")
			return false
		}
	case <-time.After(t):
		{
			println("timeout occur")
			return true
		}
	}
}
