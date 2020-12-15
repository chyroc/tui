package internal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func listenForResize(output *os.File, f func(width, height int, err error)) {
	w, h, err := terminal.GetSize(int(output.Fd()))
	f(w, h, err)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)
	for {
		<-sig
		w, h, err := terminal.GetSize(int(output.Fd()))
		fmt.Println("resize", w, h, err)
		f(w, h, err)
	}
}
