package main

import (
	"fmt"
	"time"

	"github.com/chyroc/tui"
)

type worker struct {
	msg chan string
}

func (r *worker) Init() error {
	go func() {
		for {
			r.msg <- fmt.Sprintf("time - %d", time.Now().UnixNano())
		}
	}()
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	var msg string
	select {
	case x := <-r.msg:
		msg = x
	default:
		msg = "default"
	}
	return msg
}

func NewWorker() *worker {
	return &worker{
		msg: make(chan string),
	}
}

func main() {
	r := tui.New(NewWorker())

	if err := r.Run(); err != nil {
		panic(err)
	}
}
