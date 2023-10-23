// Implementation of Timsort algorithm
// Reference: https://en.wikipedia.org/wiki/Timsort

package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
)

const runSizeThreshold = 8

// Timsort is a simple generic implementation of Timsort algorithm.
func Timsort[T constraints.Ordered](data []T) []T {
	runSize := calculateRunSize(len(data))
	insertionSortRuns(data, runSize)
	mergeRuns(data, runSize)
	return data
}

// calculateRunSize returns a run size parameter that is further used
// to slice the data slice.
func calculateRunSize(dataLength int) int {
	remainder := 0
	for dataLength >= runSizeThreshold {
		if dataLength%2 == 1 {
			remainder = 1
		}

		dataLength = dataLength / 2
	}

	return dataLength + remainder
}

// insertionSortRuns runs insertion sort on all the data runs one by one.
func insertionSortRuns[T constraints.Ordered](data []T, runSize int) {
	for lower := 0; lower < len(data); lower += runSize {
		upper := lower + runSize
		if upper >= len(data) {
			upper = len(data)
		}

		insertionSortRun(data[lower:upper])
	}
}

// insertionSortRuns runs insertion sort on a single data run slice.
func insertionSortRun[T constraints.Ordered](data []T) {
	for i := 1; i < len(data); i++ {
		value := data[i]
		j := i
		// return to the sorted part of slice by decrementing the j index and stop
		// upon reaching a smaller value
		for ; j > 0 && data[j-1] > value; j-- {
			data[j] = data[j-1]
		}

		data[j] = value
	}
}

// mergeRuns merge sorts all the data runs into a single sorted data slice.
func mergeRuns[T constraints.Ordered](data []T, runSize int) {
	for size := runSize; size < len(data); size *= 2 {
		for lowerBound := 0; lowerBound < len(data); lowerBound += size * 2 {
			middleBound := lowerBound + size - 1
			upperBound := lowerBound + 2*size - 1
			if len(data)-1 < upperBound {
				upperBound = len(data) - 1
			}

			mergeRun(data, lowerBound, middleBound, upperBound)
		}
	}
}

// mergeRun uses merge sort to sort adjacent data runs.
func mergeRun[T constraints.Ordered](data []T, lower, mid, upper int) {
	left := make([]T, mid-lower+1)
	right := make([]T, upper-mid)
	copy(left, data[lower:mid+1])
	copy(right, data[mid+1:upper+1])
	i, j, k := 0, 0, lower
	// checks the top of left and right slice, chooses the smallest value
	// and increments proper slice index until one reaches the slice's length
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}

		k++
	}

	// completes the merge sort with left-over values from left slice
	for i < len(left) {
		data[k] = left[i]
		k++
		i++
	}

	// completes the merge sort with left-over values from right slice
	for j < len(right) {
		data[k] = right[j]
		k++
		j++
	}
}
