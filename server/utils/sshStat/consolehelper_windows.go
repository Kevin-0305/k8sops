package sshStat

import (
	"io"
	"os"
	"os/exec"

	"github.com/mattn/go-colorable"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getOutput() io.Writer {
	return colorable.NewColorableStdout()
}
