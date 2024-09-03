package sortx

// HeapSort is a comparison-based sorting algorithm that uses a binary heap data structure.
// It first builds a max heap from the input array, then repeatedly extracts the maximum element from the heap and rebuilds the heap.
// Target Time Complexity: O(n log n)
// Worst Time Complexity: O(n log n)
// Best Time Complexity: O(n log n)
// Space Complexity: O(1)
func HeapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest element among the root, left child, and right child
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root, swap it with the root and heapify the affected subtree
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
