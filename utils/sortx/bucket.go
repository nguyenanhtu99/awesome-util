package sortx

// Space Complexity: O(n).
func BucketSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find the maximum and minimum values in the input array
	minVal, maxVal := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		} else if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	// Create buckets
	bucketSize := 10
	bucketCount := (maxVal-minVal)/bucketSize + 1
	buckets := make([][]int, bucketCount)
	for i := range bucketCount {
		buckets[i] = make([]int, 0)
	}

	// Distribute elements into buckets
	for i := range n {
		bucketIndex := (arr[i] - minVal) / bucketSize
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}

	// Sort each bucket and update the input array
	index := 0
	for i := range bucketCount {
		InsertionSort(buckets[i])
		for j := range len(buckets[i]) {
			arr[index] = buckets[i][j]
			index++
		}
	}
}
