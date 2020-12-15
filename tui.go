package tui

type TUI interface {
	Run() error
}

func New() TUI {
	return newImpl()
}
