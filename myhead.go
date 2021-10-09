package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ashishkujoy/go-tools/utils"
	"io"
	"os"
)

func main() {
	n := flag.Int("n", 10, "number of lines")
	flag.Parse()
	args := flag.Args()

	var files []*os.File

	if len(args) == 0 {
		files = []*os.File{os.Stdin}
	} else {
		for _, file := range args {
			fileToReadFrom, err := os.Open(file)
			utils.ExitIfErrNotNil(err)
			files = append(files, fileToReadFrom)
		}
	}

	filesCount := len(files)

	for _, file := range files {
		if filesCount > 1 {
			_, err := io.WriteString(os.Stdout, fmt.Sprintf("\n\n==================== %s ===================== \n\n", file.Name()))
			utils.ExitIfErrNotNil(err)
		}
		scanner := bufio.NewScanner(file)

		linesRead := 0

		for linesRead < *n && scanner.Scan() {
			_, err := io.WriteString(os.Stdout, scanner.Text())
			utils.ExitIfErrNotNil(err)
			_, err = io.WriteString(os.Stdout, "\n")
			utils.ExitIfErrNotNil(err)
			linesRead++
		}
	}

	os.Exit(0)
}
