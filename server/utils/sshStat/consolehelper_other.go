package sshStat

import (
	"io"
	"os"
)

func clearConsole() {}

func getOutput() io.Writer {
	return os.Stdout
}
