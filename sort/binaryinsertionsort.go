// Binary Insertion Sort
// description: Implementation of binary insertion sort in Go
// details: Binary Insertion Sort is a variation of
// Insertion sort in which proper location to
// insert the selected element is found using the
// Binary search algorithm.
// ref: https://www.geeksforgeeks.org/binary-insertion-sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

func BinaryInsertion[T constraints.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		low := 0
		high := currentIndex - 1

		for low <= high {
			mid := low + (high-low)/2
			if arr[mid] > temporary {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		for itr := currentIndex; itr > low; itr-- {
			arr[itr] = arr[itr-1]
		}

		arr[low] = temporary
	}
	return arr
}
