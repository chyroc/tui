package internal

import (
	"fmt"
	"io"

	"github.com/muesli/termenv"
)

func HideCursor(w io.Writer) {
	TriggerConsoleCSI(w, termenv.HideCursorSeq)
}

func ShowCursor(w io.Writer) {
	TriggerConsoleCSI(w, termenv.ShowCursorSeq)
}

func TriggerConsoleCSI(w io.Writer, seq string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, termenv.CSI+seq, a...)
}

func clearLine(w io.Writer) {
	TriggerConsoleCSI(w, termenv.EraseLineSeq, 2)
}

func cursorUp(w io.Writer) {
	TriggerConsoleCSI(w, termenv.CursorUpSeq, 1)
}

func cursorDown(w io.Writer) {
	TriggerConsoleCSI(w, termenv.CursorDownSeq, 1)
}

func insertLine(w io.Writer, numLines int) {
	TriggerConsoleCSI(w, "%dL", numLines)
}

func moveCursor(w io.Writer, row, col int) {
	TriggerConsoleCSI(w, termenv.CursorPositionSeq, row, col)
}

func changeScrollingRegion(w io.Writer, top, bottom int) {
	TriggerConsoleCSI(w, termenv.ChangeScrollingRegionSeq, top, bottom)
}

func cursorBack(w io.Writer, n int) {
	TriggerConsoleCSI(w, termenv.CursorBackSeq, n)
}
