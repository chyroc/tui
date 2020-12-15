package tui

type TUI interface {
	Run() error
}

type Worker interface {
	Init() error
	Close() error
	View() string
}

func New(worker Worker) TUI {
	return newImpl(worker)
}
