package sortx

// Space Complexity: O(n).
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

//nolint:gocritic // ignore
func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays
	for i := range n1 {
		L[i] = arr[l+i]
	}
	for j := range n2 {
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
