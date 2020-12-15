package main

import (
	"github.com/chyroc/tui"
)

type worker struct {
}

func (r *worker) Init() error {
	return nil
}

func (r *worker) Close() error {
	return nil
}

func NewWorker() *worker {
	return &worker{}
}

func main() {
	r := tui.New(NewWorker())

	if err := r.Run(); err != nil {
		panic(err)
	}
}
