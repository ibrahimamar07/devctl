package runner

import (
	"os"
	"os/exec"
)

func Run(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Dir, _ = os.Getwd()
	return cmd.Run()
}
