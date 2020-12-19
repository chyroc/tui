package tui

import (
	"github.com/tj/go-terminput"
)

type TUI interface {
	Run() error
	Stop()
	SetWorker(worker Worker)
}

type Worker interface {
	Init() error
	Close() error
	View() string
	HandleInput(e *terminput.KeyboardInput)
}

func New() TUI {
	return newImpl()
}
