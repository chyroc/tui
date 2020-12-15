package internal

import (
	"fmt"
	"io"

	"github.com/muesli/termenv"
)

func HideCursor(w io.Writer) {
	triggerConsoleCSI(w, termenv.HideCursorSeq)
}

func ShowCursor(w io.Writer) {
	triggerConsoleCSI(w, termenv.ShowCursorSeq)
}

func triggerConsoleCSI(w io.Writer, seq string) {
	_, _ = fmt.Fprintf(w, termenv.CSI+seq)
}
