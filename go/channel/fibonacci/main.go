package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y <= math.MaxInt64; c <- y {
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	for x, ok := <-c; ok; x, ok = <-c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
