package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"os"
	"regexp"
	"sync"
)

func main() {
	wordToCount := flag.String("w", "", "word to count")
	flag.Parse()
	filePaths := flag.Args()[:]
	wordCountInputChannel := make(chan int)
	wordCountOutputChannel := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(filePaths) + 1)

	regexpToMatch, err := regexp.Compile(*wordToCount)
	utils.ExitIfErrNotNil(err)

	go func() {
		count := add(wordCountOutputChannel)
		fmt.Printf("Total Count: %d\n", count)
		waitGroup.Done()
	}()

	for _, filePath := range filePaths {
		go func(filePath string) {
			wordCountInputChannel <- countWordOccurrence(filePath, regexpToMatch)
			waitGroup.Done()
		}(filePath)
	}

	take(wordCountOutputChannel, wordCountInputChannel, len(filePaths))
	close(wordCountInputChannel)

	waitGroup.Wait()
}

func countWordOccurrence(filePath string, word *regexp.Regexp) int {
	wordCount := 0
	file, err := os.Open(filePath)

	utils.ExitIfErrNotNil(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if word.MatchString(scanner.Text()) {
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
