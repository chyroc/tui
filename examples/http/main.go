package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chyroc/tui"
)

type worker struct {
	msg string
	tui tui.TUI
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
		tui.Stop(r.tui)
	}()
	return nil
}

func (r *worker) Close() error {
	return nil
}

func (r *worker) View() string {
	return r.msg
}

func NewWorker() *worker {
	return &worker{
	}
}

func main() {
	worker := NewWorker()
	r := tui.New(worker)
	worker.tui = r

	if err := r.Run(); err != nil {
		panic(err)
	}
}
