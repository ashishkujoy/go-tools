package main

import (
	"bufio"
	"github.com/ashishkujoy/go-tools/utils"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		utils.ExitIfErrNotNil(err)
		os.Exit(0)
	}

	filename := args[2]

	file, err := os.Open(filename)
	utils.ExitIfErrNotNil(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_, err := io.WriteString(os.Stdout, scanner.Text())
		utils.ExitIfErrNotNil(err)
		_, _ = io.WriteString(os.Stdout, "\n")
	}

	os.Exit(0)
}
