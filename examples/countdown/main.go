package main

import (
	"fmt"
	"time"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui"
)

type worker struct {
	tui tui.TUI
	end time.Time
}

func (r *worker) Init() error {
	go func() {
		end := time.Now().Add(time.Second * 5)
		r.end = end
		for {
			if time.Now().After(end) {
				break
			}
			time.Sleep(time.Second / 10)
		}
		r.tui.Stop()
	}()
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	return fmt.Sprintf("will stop after %0.2f", r.end.Sub(time.Now()).Seconds())
}

func (r *worker) HandleInput(e *terminput.KeyboardInput) {
	if e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == tui.KeyCtrlC {
		r.tui.Stop()
		return
	}
	fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
}

func NewWorker(tui tui.TUI) *worker {
	return &worker{
		tui: tui,
	}
}

func main() {
	worker := NewWorker(tui.New())
	worker.tui.SetWorker(worker)

	if err := worker.tui.Run(); err != nil {
		panic(err)
	}
}
