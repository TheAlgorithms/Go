package sort

import "github.com/TheAlgorithms/Go/constraints"

// Bucket sorts a slice. It is mainly useful
// when input is uniformly distributed over a range.
func Bucket[T constraints.Number](arr []T) []T {
	// early return if the array too small
	if len(arr) <= 1 {
		return arr
	}

	// find the maximum and minimum elements in arr
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	// create an empty bucket for each element in arr
	bucket := make([][]T, len(arr))

	// put each element in the appropriate bucket
	for _, v := range arr {
		bucketIndex := int((v - min) / (max - min) * T(len(arr)-1))
		bucket[bucketIndex] = append(bucket[bucketIndex], v)
	}

	// use insertion sort to sort each bucket
	for i := range bucket {
		bucket[i] = Insertion(bucket[i])
	}

	// concatenate the sorted buckets
	sorted := make([]T, 0, len(arr))
	for _, v := range bucket {
		sorted = append(sorted, v...)
	}

	return sorted
}
