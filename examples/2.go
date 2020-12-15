package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chyroc/tui/internal"
)

var i int

func randInput() string {
	i++
	rand.Seed(time.Now().UnixNano())
	s := fmt.Sprintf("= %d =\n", i)
	for i := 0; i < int(rand.Int31n(10)); i++ {
		if i > 0 {
			s += "\n"
		}

		s += strings.Repeat("x", int(rand.Int31n(20)))
	}
	return s
}

func main() {
	r := internal.NewRenderer(os.Stdout, new(sync.Mutex))

	r.Start()
	defer r.Stop()

	for {
		r.Write(randInput())
		time.Sleep(time.Second)
	}
}
