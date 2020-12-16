package main

import (
	"fmt"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui"
)

type worker struct {
	options  []string
	selected int
	tui      tui.TUI
}

func (r *worker) Init() error {
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	s := "Select one hero?\n\n"
	for idx, v := range r.options {
		if idx == r.selected {
			s += "(•) " + v
		} else {
			s += "( ) " + v
		}
		s += "\n"
	}
	s += "\n"
	return s
}

func (r *worker) HandleInput(e *terminput.KeyboardInput) {
	switch {
	case e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == tui.KeyCtrlC:
		tui.Stop(r.tui)
	case e.Key() == terminput.KeyUp:
		if r.selected > 0 {
			r.selected--
		}
	case e.Key() == terminput.KeyDown:
		if r.selected < len(r.options)-1 {
			r.selected++
		}
	case e.Key() == terminput.KeyEnter:
		fmt.Println("select:", r.options[r.selected])
		tui.Stop(r.tui)
	default:
		fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
	}
}

func NewWorker() *worker {
	return &worker{}
}

func main() {
	worker := NewWorker()
	r := tui.New(worker)
	worker.tui = r
	worker.options = []string{
		"齐木楠雄",
		"铁臂阿童木",
		"孙悟空",
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
