package main

import (
	"github.com/ashishkujoy/go-tools/utils"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	path, err := exec.LookPath("ps")
	utils.ExitIfErrNotNil(err)

	env := os.Environ()

	command := []string{"ps", "-a", "-x"}
	err = syscall.Exec(path, command, env)


	utils.ExitIfErrNotNil(err)
}
