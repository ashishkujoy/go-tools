package main

import (
	"fmt"
	"sync"
)

func main() {
	numberIn := make(chan int)
	squareOut := make(chan int)
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	go genNumbers(0, 100, numberIn, &waitGroup)
	go square(numberIn, squareOut, &waitGroup)
	go sum(squareOut, &waitGroup)
	waitGroup.Wait()
}

func genNumbers(min, max int, out chan<- int, group *sync.WaitGroup) {
	for i := min; i < max; i++ {
		out <- i
	}
	close(out)
	group.Done()
}

func square(in <-chan int, out chan<- int, group *sync.WaitGroup) {
	for n := range in {
		out <- n * n
	}
	close(out)
	group.Done()
}

func sum(in <-chan int, group *sync.WaitGroup) {
	sum := 0
	for n := range in {
		sum += n
	}
	fmt.Printf("Sum of squares %d\n", sum)
	group.Done()
}
