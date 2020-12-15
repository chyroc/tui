package tui

import (
	"fmt"
	"os"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui/internal"
)

type impl struct {
	output *os.File
	worker Worker
}

func newImpl(worker Worker) *impl {
	return &impl{
		output: os.Stdout,
		worker: worker,
	}
}

func (r *impl) Run() error {
	if err := initConsole(r.output); err != nil {
		return err
	}
	defer resetConsole(r.output)

	if err := r.worker.Init(); err != nil {
		return err
	}
	defer r.worker.Close()

	fmt.Println("start read")

	// read input
	for {
		e, err := internal.ReadTerminal(r.output)
		if err != nil {
			return err
		}

		if e.Key() == terminput.KeyEscape || e.Rune() == 'q' {
			break
		}

		fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v, ctrl-a=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key(), e.Key() == internal.KeyCtrlA)
	}

	return nil
}
