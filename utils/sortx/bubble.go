package sortx

// BubbleSort is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
// The pass through the list is repeated until the list is sorted.
// Target Time Complexity: O(n^2)
// Worst Time Complexity: O(n^2)
// Best Time Complexity: O(n)
// Space Complexity: O(1)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}