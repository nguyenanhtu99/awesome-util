package sortx

// BucketSort is a sorting algorithm that works by distributing the elements of an array into a number of buckets.
// Each bucket is then sorted individually, either using a different sorting algorithm, or by recursively applying the bucket sorting algorithm.
// Target Time Complexity: O(n + k)
// Worst Time Complexity: O(n^2)
// Best Time Complexity: O(n)
// Space Complexity: O(n)
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
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// Distribute elements into buckets
	for i := 0; i < n; i++ {
		bucketIndex := (arr[i] - minVal) / bucketSize
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}

	// Sort each bucket and update the input array
	index := 0
	for i := 0; i < bucketCount; i++ {
		InsertionSort(buckets[i])
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
}
