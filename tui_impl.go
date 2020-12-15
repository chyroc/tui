package tui

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/tj/go-terminput"

	"github.com/chyroc/tui/internal"
)

type impl struct {
	output *os.File
	worker Worker
	render *internal.Renderer
	end    chan int
}

func newImpl(worker Worker) *impl {
	lock := sync.Mutex{}
	return &impl{
		output: os.Stdout,
		worker: worker,
		render: internal.NewRenderer(os.Stdout, &lock),
		end:    make(chan int),
	}
}

func (r *impl) Run() (finalErr error) {
	// console
	if err := initConsole(r.output); err != nil {
		return err
	}
	defer resetConsole(r.output)

	// render
	r.render.Start()
	defer r.render.Stop()

	// worker
	if err := r.worker.Init(); err != nil {
		return err
	}
	defer r.worker.Close()

	// read input
	go func() {
		for {
			e, err := internal.ReadTerminal(r.output)
			if err != nil {
				finalErr = err
				return
			}

			if e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == internal.KeyCtrlC {
				close(r.end)
				break
			}

			fmt.Printf("e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v, ctrl-a=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key(), e.Key() == internal.KeyCtrlA)
		}
	}()

	t := time.NewTicker(time.Second / 60)
	for {
		select {
		case <-t.C:
			v := r.worker.View()
			r.render.Write(v)
		case <-r.end:
			return nil
		}
	}

	return nil
}
