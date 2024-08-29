package sortx

// QuickSort is a sorting algorithm that uses divide and conquer strategy to sort an array.
// It picks an element as pivot and partitions the given array around the picked pivot.
// There are many different versions of quickSort that pick pivot in different ways.
// The key process in quickSort is partition().
// Target Time Complexity: O(n log n)
// Worst Time Complexity: O(n^2)
// Best Time Complexity: O(n log n)
// Space Complexity: O(log n)
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivotIndex := partition(arr)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

// partition function partitions the array around the pivot
func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
