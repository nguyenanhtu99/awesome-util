package sortx

// Space Complexity: O(log n).
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivotIndex := partition(arr)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

// partition function partitions the array around the pivot.
func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := range len(arr) - 1 {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
