package sortx

// SelectionSort is a sorting algorithm that divides the input list into two parts: the sublist of items already sorted and the sublist of items remaining to be sorted.
// It repeatedly finds the minimum element from the unsorted part and swaps it with the first element of the unsorted part.
// Target Time Complexity: O(n^2)
// Worst Time Complexity: O(n^2)
// Best Time Complexity: O(n^2)
// Space Complexity: O(1)
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}