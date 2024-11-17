package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	doubleFunc := func(i int) int { return i * 2 }

	for i, num := range IteratorTransform(list, doubleFunc) {
		fmt.Println(i, num)
	}
}
