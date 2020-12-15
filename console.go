package tui

import (
	"io"

	"github.com/containerd/console"

	"github.com/chyroc/tui/internal"
)

var consoleIns console.Console

func initConsole(w io.Writer) error {
	consoleIns = console.Current()
	err := consoleIns.SetRaw()
	if err != nil {
		return err
	}

	internal.HideCursor(w)

	return nil
}

func resetConsole(w io.Writer) error {
	internal.ShowCursor(w)
	return consoleIns.Reset()
}
