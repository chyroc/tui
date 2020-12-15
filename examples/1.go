package main

import (
	"github.com/chyroc/tui"
)

func main() {
	r := tui.New()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
