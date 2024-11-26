// quicksort.go
// description: Implementation of in-place quicksort algorithm
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
// worst time complexity: O(n^2)
// average time complexity: O(n log n)
// space complexity: O(log n)
// author(s) [Taj](https://github.com/tjgurwara99)
// see sort_test.go for a test implementation, test function TestQuickSort.

package sort

import "github.com/TheAlgorithms/Go/constraints"

func Partition[T constraints.Ordered](arr []T, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

// QuicksortRange Sorts the specified range within the array
func QuicksortRange[T constraints.Ordered](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := Partition(arr, low, high)
		QuicksortRange(arr, low, pivot-1)
		QuicksortRange(arr, pivot+1, high)
	}
}

// Quicksort Sorts the entire array
func Quicksort[T constraints.Ordered](arr []T) []T {
	QuicksortRange(arr, 0, len(arr)-1)
	return arr
}
