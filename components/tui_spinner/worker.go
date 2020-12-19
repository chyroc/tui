package tui_spinner

import (
	"fmt"
	"time"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui"
)

type worker struct {
	tui      tui.TUI
	count    int
	options  [][]string
	selected int
	fps      time.Duration
}

func (r *worker) Init() error {
	if r.fps == 0 {
		r.fps = time.Second / 10 // 一秒10变
	}
	go func() {
		for {
			r.count++
			time.Sleep(r.fps)
		}
	}()
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	count := r.count
	options := r.options
	selected := r.selected

	// fmt.Println("view", count, selected, options[selected][count%len(options[selected])])
	return options[selected][count%len(options[selected])]
}

func (r *worker) HandleInput(e *terminput.KeyboardInput) {
	switch {
	case e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == tui.KeyCtrlC:
		r.tui.Stop()
	case e.Key() == terminput.KeyUp || e.Key() == terminput.KeyLeft:
		if r.selected > 0 {
			r.selected--
		} else {
			r.selected = len(r.options) - 1
		}
		// len(r.options[r.selected])
	case e.Key() == terminput.KeyDown || e.Key() == terminput.KeyRight:
		if r.selected < len(r.options)-1 {
			r.selected++
		} else {
			r.selected = 0
		}
	default:
		fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
	}
}
