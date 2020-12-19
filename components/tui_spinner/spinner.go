package tui_spinner

import (
	"github.com/chyroc/tui"
)

func New() *ComponentsSpinner {
	worker := &worker{
		tui: tui.New(),
	}
	worker.tui.SetWorker(worker)

	return &ComponentsSpinner{
		worker: worker,
	}
}

type ComponentsSpinner struct {
	worker *worker
}

// Run 运行 tui.Run
func (r *ComponentsSpinner) Run() error {
	err := r.worker.tui.Run()
	if err != nil {
		return err
	}
	// if r.worker.err != nil {
	// 	return 0, "", r.worker.err
	// }
	return nil
}
