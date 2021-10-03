package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	sourceFilePath := os.Args[1]
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

	fmt.Printf("Line: %v, Word: %v, Characters: %v\n", lineCount, wordCount, characterCount)
}
