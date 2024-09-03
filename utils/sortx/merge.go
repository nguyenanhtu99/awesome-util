package sortx

// MergeSort is a divide-and-conquer algorithm that divides the input list into two halves, recursively sorts the two halves, and then merges the sorted halves.
// Target Time Complexity: O(n log n)
// Worst Time Complexity: O(n log n)
// Best Time Complexity: O(n log n)
// Space Complexity: O(n)
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temporary arrays
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[] if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[] if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
