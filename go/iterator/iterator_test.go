package main

import (
	"reflect"
	"testing"
)

var (
	transform = func(i int) int { return i * 2 }
	list      = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func TestTransform(t *testing.T) {
	// Test case 1: Transform integers to their squares
	intList := []int{1, 2, 3, 4, 5}
	expectedSquares := []int{1, 4, 9, 16, 25}
	transformFunc := func(x int) int { return x * x }

	normalSquares := NormalTransform(intList, transformFunc)
	for i, v := range normalSquares {
		if v != expectedSquares[i] {
			t.Errorf(
				"Normal: Test failed for integers to squares. Index %d: got %v, expected %v",
				i,
				v,
				expectedSquares[i],
			)
		}
	}

	iteratorSquares := IteratorTransform(intList, transformFunc)
	for i, v := range iteratorSquares {
		if v != expectedSquares[i] {
			t.Errorf(
				"Iterator: Test failed for integers to squares. Index %d: got %v, expected %v",
				i,
				v,
				expectedSquares[i],
			)
		}
	}

	// Test case 2: Transform strings to their lengths
	stringList := []string{"apple", "banana", "kiwi"}
	expectedLengths := []int{5, 6, 4}
	lengthFunc := func(s string) int { return len(s) }

	normalLengths := NormalTransform(stringList, lengthFunc)
	for i, v := range normalLengths {
		if v != expectedLengths[i] {
			t.Errorf(
				"Normal: Test failed for strings to lengths. Index %d: got %v, expected %v",
				i,
				v,
				expectedLengths[i],
			)
		}
	}

	iteratorLengths := IteratorTransform(stringList, lengthFunc)
	for i, v := range iteratorLengths {
		if v != expectedLengths[i] {
			t.Errorf(
				"Iterator: Test failed for strings to lengths. Index %d: got %v, expected %v",
				i,
				v,
				expectedLengths[i],
			)
		}
	}

	// Test case 3: Transform floats to their negation
	floatList := []float64{1.1, 2.2, 3.3}
	expectedNegations := []float64{-1.1, -2.2, -3.3}
	negationFunc := func(f float64) float64 { return -f }

	normalNegations := NormalTransform(floatList, negationFunc)
	for i, v := range normalNegations {
		if !reflect.DeepEqual(v, expectedNegations[i]) {
			t.Errorf(
				"Normal: Test failed for floats to negations. Index %d: got %v, expected %v",
				i,
				v,
				expectedNegations[i],
			)
		}
	}

	iteratorNegations := IteratorTransform(floatList, negationFunc)
	for i, v := range iteratorNegations {
		if !reflect.DeepEqual(v, expectedNegations[i]) {
			t.Errorf(
				"Iterator: Test failed for floats to negations. Index %d: got %v, expected %v",
				i,
				v,
				expectedNegations[i],
			)
		}
	}
}

func BenchmarkNormalTransform(b *testing.B) {
	for range b.N {
		for _, num := range NormalTransform(list, transform) {
			if num == 4 {
				break
			}
		}
	}
}

func BenchmarkIteratorTransform(b *testing.B) {
	for range b.N {
		for _, num := range IteratorTransform(list, transform) {
			if num == 4 {
				break
			}
		}
	}
}
