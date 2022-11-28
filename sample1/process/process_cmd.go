package process

import (
	"os"
	"os/exec"
)

func NewCmd(command string) {
	println(command, "begin start")
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
