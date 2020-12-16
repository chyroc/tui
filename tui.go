package tui

import (
	"github.com/tj/go-terminput"
)

type TUI interface {
	Run() error
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
	close(r.end)
}

func New(worker Worker) TUI {
	return newImpl(worker)
}
