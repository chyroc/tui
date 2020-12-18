package tui_select

import (
	"errors"
	"fmt"

	"github.com/chyroc/tui"
)

var ErrNoSelect = errors.New("no select err")

func NewStringSelect() *ComponentsSelect {
	worker := &worker{
		tui: tui.New(),
	}
	worker.tui.SetWorker(worker)

	return &ComponentsSelect{
		worker: worker,
	}
}

type ComponentsSelect struct {
	worker *worker
}

func (r *ComponentsSelect) SetTitle(title string) {
	r.worker.title = title
}

func (r *ComponentsSelect) SetOptions(options []string) {
	r.worker.options = options
}

func (r *ComponentsSelect) SetSize(size int) {
	if size <= 0 {
		panic(fmt.Sprintf("invalid size: %d", size))
	}

	r.worker.size = size
	r.worker.windowMin = 0
	r.worker.windowMax = size - 1
}

// 运行 tui.Run
func (r *ComponentsSelect) Select() (int, string, error) {
	err := r.worker.tui.Run()
	if err != nil {
		return 0, "", err
	}
	if r.worker.err != nil {
		return 0, "", r.worker.err
	}
	return r.worker.selected, r.worker.options[r.worker.selected], nil
}
