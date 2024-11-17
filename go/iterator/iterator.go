package main

import (
	"context"
	"iter"
	"sync"
)

func NormalTransform[T1, T2 any](list []T1, transform func(T1) T2) []T2 {
	transformed := make([]T2, len(list))

	for i, t := range list {
		transformed[i] = transform(t)
	}

	return transformed
}

func IteratorTransform[T1, T2 any](list []T1, transform func(T1) T2) iter.Seq2[int, T2] {
	return func(yield func(int, T2) bool) {
		for i, t := range list {
			if !yield(i, transform(t)) {
				return
			}
		}
	}
}

func Parallel[T any](list []T) func(func(i int, t T) bool) {
	return func(yield func(int, T) bool) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(len(list))

		for i, t := range list {
			go func(i int, t T) {
				defer wg.Done()

				select {
				case <-ctx.Done():
					return
				default:
					if !yield(i, t) {
						cancel()
					}
				}

			}(i, t)
		}
		wg.Wait()
	}
}
