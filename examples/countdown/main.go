package main

import (
	"fmt"
	"time"

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
