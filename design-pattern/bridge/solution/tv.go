package main

import (
	"fmt"

	"github.com/samber/lo"
)

type tv struct {
	volume int
}

func (t *tv) setVolume(volume int) {
	t.volume = lo.Clamp(volume, 0, 99)

	fmt.Printf("The TV volume is %d\n", t.volume)
}

func (t *tv) getVolume() int {
	return t.volume
}
