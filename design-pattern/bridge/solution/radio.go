package main

import (
	"fmt"

	"github.com/samber/lo"
)

type radio struct {
	volume int
}

func (r *radio) setVolume(volume int) {
	r.volume = lo.Clamp(volume, 0, 100)

	fmt.Printf("The radio volume is %d\n", r.volume)
}

func (r *radio) getVolume() int {
	return r.volume
}
