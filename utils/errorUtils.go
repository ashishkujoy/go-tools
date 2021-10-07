package utils

import (
	"fmt"
	"os"
)

func ExitIfErrNotNil(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
