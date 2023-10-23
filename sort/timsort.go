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

func insertionSortRuns[T constraints.Ordered](data []T, runSize int) {
	for lower := 0; lower < len(data); lower += runSize {
		upper := lower + runSize
		if upper >= len(data) {
			upper = len(data)
		}

		insertionSortRun(data[lower:upper])
	}
}

func insertionSortRun[T constraints.Ordered](data []T) {
	for i := 1; i < len(data); i++ {
		value := data[i]
		j := i
		for ; j > 0 && data[j-1] > temp; j-- {
		for ; j > 0 && data[j-1] > value; j-- {
			data[j] = data[j-1]
		}

		data[j] = value
	}
}

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

func mergeRun[T constraints.Ordered](data []T, lower, mid, upper int) {
	left := make([]T, mid-lower+1)
	right := make([]T, upper-mid)
	copy(left, data[lower:mid+1])
	copy(right, data[mid+1:upper+1])
	i, j, k := 0, 0, lower
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

	for i < len(left) {
		data[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		data[k] = right[j]
		k++
		j++
	}
}
