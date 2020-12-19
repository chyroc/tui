package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui"
)

type worker struct {
	tui.TUI
	msg string
}

func (r *worker) Init() error {
	go func() {
		url := "http://example.org/"
		r.msg = fmt.Sprintf("start fetch %s ...", url)
		time.Sleep(time.Second)
		resp, err := http.Get(url)
		if err != nil {
			r.msg = fmt.Sprintf("failed: %s", err)
		} else {
			r.msg = fmt.Sprintf("success, code: %d", resp.StatusCode)
		}
		time.Sleep(time.Second)
		r.TUI.Stop()
	}()
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	return r.msg
}

func (r *worker) HandleInput(e *terminput.KeyboardInput) {
	if e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == tui.KeyCtrlC {
		r.TUI.Stop()
		return
	}
	fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
}

func NewWorker(tui tui.TUI) *worker {
	return &worker{TUI: tui}
}

func main() {
	worker := NewWorker(tui.New())
	worker.SetWorker(worker)

	if err := worker.Run(); err != nil {
		panic(err)
	}
}
