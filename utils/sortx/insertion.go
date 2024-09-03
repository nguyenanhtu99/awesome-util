package sortx

// InsertionSort is a simple sorting algorithm that builds the final sorted array one item at a time.
// It takes each element from the input list and inserts it into its correct position in the sorted list.
// Target Time Complexity: O(n^2)
// Worst Time Complexity: O(n^2)
// Best Time Complexity: O(n)
// Space Complexity: O(1)
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}