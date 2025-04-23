package sortx

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Array with one element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Array with multiple elements",
			input:    []int{9, 3, 7, 1, 5},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Array with multiple elements (sorted)",
			input:    []int{1, 3, 5, 7, 9},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Array with multiple elements (reverse sorted)",
			input:    []int{9, 7, 5, 3, 1},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Array with multiple elements (all equal)",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "Array with multiple elements (negative numbers)",
			input:    []int{-9, -3, -7, -1, -5},
			expected: []int{-9, -7, -5, -3, -1},
		},
		// Add more test cases here...
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			HeapSort(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, test.input)
			}
		})
	}
}
