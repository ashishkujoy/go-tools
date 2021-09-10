package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	source := os.Args[1]
	destination := os.Args[2]

	err := filepath.Walk(source, walkFunc(source, destination))
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func walkFunc(source, destination string) filepath.WalkFunc {
	return func(path string, stat os.FileInfo, err error) error {
		dest := newPath(source, destination, path)
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil
		}

		if fileInfo.IsDir() {
			return os.MkdirAll(dest, fileInfo.Mode())
		} else {
			return copy(path, dest, 1024)
		}
	}
}

func newPath(source, destination, currentPath string) string {
	return strings.Replace(currentPath, source, destination, 1)
}

func copy(source, destination string, bufferSize int) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	sourceStat, err := os.Stat(source)

	if err != nil {
		return err
	}

	destinationFile, err := os.Create(destination)

	if err != nil {
		return err
	}

	err = os.Chmod(destination, sourceStat.Mode())

	if err != nil {
		return err
	}

	buf := make([]byte, bufferSize)

	for {
		numberOfBytes, err := sourceFile.Read(buf)

		if err != nil && err != io.EOF {
			return err
		}

		if numberOfBytes == 0 {
			break
		}

		if _, err = destinationFile.Write(buf[:numberOfBytes]); err != nil {
			return nil
		}
	}
	return err
}
