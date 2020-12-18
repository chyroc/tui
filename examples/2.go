package main

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

func main() {
	fmt.Print(termenv.CSI + "5m")
	time.Sleep(time.Hour)
}
