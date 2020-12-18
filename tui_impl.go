package tui

import (
	"os"
	"sync"
	"time"

	"github.com/chyroc/tui/internal"
)

type impl struct {
	output    *os.File
	worker    Worker
	render    *internal.Renderer
	done      chan struct{}
	closeOnce sync.Once
}

func newImpl() *impl {
	lock := sync.Mutex{}
	return &impl{
		output: os.Stdout,
		render: internal.NewRenderer(os.Stdout, &lock),
		done:   make(chan struct{}),
	}
}

func (r *impl) SetWorker(worker Worker) {
	r.worker = worker
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
			select {
			case <-r.done:
				return
			default:
				e, err := readTerminalInput(r.output)
				if err != nil {
					finalErr = err
					return
				}

				r.worker.HandleInput(e)
			}
		}
	}()

	t := time.NewTicker(time.Second / 60)
	for {
		select {
		case <-t.C:
			v := r.worker.View()
			r.render.Write(v)
		case <-r.done:
			return nil
		}
	}
}
