package tui

import (
	"github.com/tj/go-terminput"

	"github.com/chyroc/tui/internal"
)

type TUI interface {
	Run() error
	SetWorker(worker Worker)
}

type Worker interface {
	Init() error
	Close() error
	View() string
	HandleInput(e *terminput.KeyboardInput)
}

func Stop(t TUI) {
	if t == nil {
		return
	}
	r := t.(*impl)
	internal.CloseChanStruct(r.done)
}

func New() TUI {
	return newImpl()
}
