package main

import (
	"fmt"
	"github.com/bsipos/thist"
	"math/rand"
	"time"
)

// randStream return a channel filled with endless normal random values
func randStream() chan float64 {
	c := make(chan float64)
	go func() {
		for {
			c <- rand.NormFloat64()
		}
	}()
	return c
}

func main() {
	// create new histogram
	h := thist.NewHist(nil, "Example histogram", "auto", -1, true)
	c := randStream()

	i := 0
	for {
		// add data point to hsitogram
		h.Update(<-c)
		if i%50 == 0 {
			// draw histogram
			fmt.Println(h.Draw())
			time.Sleep(time.Second)
		}
		i++
	}
}
