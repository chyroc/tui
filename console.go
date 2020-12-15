package tui

import (
	"io"

	"github.com/containerd/console"
)

var consoleIns console.Console

func initConsole(w io.Writer) error {
	consoleIns = console.Current()
	err := consoleIns.SetRaw()
	if err != nil {
		return err
	}

	return nil
}

func resetConsole() error {
	return consoleIns.Reset()
}
