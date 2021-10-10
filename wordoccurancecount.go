package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"os"
	"sync"
)

func main() {
	wordToCount := flag.String("w", "", "word to count")
	flag.Parse()
	filePaths := flag.Args()[:]
	println(len(filePaths))
	wordCountInputChannel := make(chan int)
	wordCountOutputChannel := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(filePaths) + 1)

	go func() {
		count := add(wordCountOutputChannel)
		fmt.Printf("Total Count: %d\n", count)
		waitGroup.Done()
	}()

	for _, filePath := range filePaths {
		go func(filePath string) {
			wordCountInputChannel <- countWordOccurrence(filePath, *wordToCount)
			waitGroup.Done()
		}(filePath)
	}

	take(wordCountOutputChannel, wordCountInputChannel, len(filePaths))
	close(wordCountInputChannel)

	waitGroup.Wait()
}

func countWordOccurrence(filePath, word string) int {
	wordCount := 0
	file, err := os.Open(filePath)

	utils.ExitIfErrNotNil(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == word {
			wordCount++
		}
	}
	fmt.Printf("Processing done for %s:\n", filePath)
	return wordCount
}

func take(out chan<- int, in <-chan int, count int) {
	if count == 0 {
		return
	}

	taken := 0

	for element := range in {
		out <- element
		taken++
		if taken == count {
			close(out)
			return
		}
	}
}

func add(in <-chan int) int {
	total := 0
	for i := range in {
		total += i
	}
	return total
}
