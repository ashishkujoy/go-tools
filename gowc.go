package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

func main() {
	sourceFilePaths := os.Args[1:]
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(sourceFilePaths))

	for _, sourceFilePath := range sourceFilePaths {
		go func(sourceFilePath string) {
			countWords(sourceFilePath)
			waitGroup.Done()
		}(sourceFilePath)
	}

	waitGroup.Wait()
}

func countWords(sourceFilePath string) {
	sourceFile, err := os.Open(sourceFilePath)
	defer sourceFile.Close()

	if err != nil {
		fmt.Errorf("failed to read input file: %v, %v", sourceFilePath, err)
		os.Exit(2)
	}

	r := regexp.MustCompile("[^\\s]+")
	scanner := bufio.NewScanner(sourceFile)
	wordCount := 0
	lineCount := 0
	characterCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for range r.FindAllString(line, -1) {
			wordCount++
		}
		lineCount++
		characterCount += len(line)
	}

	fmt.Printf("Filename: %s Line: %v, Word: %v, Characters: %v\n", sourceFilePath, lineCount, wordCount, characterCount)
}
