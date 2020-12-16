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
		tui.Stop(r.tui)
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
		tui.Stop(r.tui)
		return
	}
	fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
}

func NewWorker() *worker {
	return &worker{}
}

func main() {
	worker := NewWorker()
	r := tui.New(worker)
	worker.tui = r

	if err := r.Run(); err != nil {
		panic(err)
	}
}
