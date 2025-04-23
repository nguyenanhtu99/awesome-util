package sortx

// Space Complexity: O(1).
func SelectionSort(arr []int) {
	n := len(arr)
	for i := range n - 1 {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
