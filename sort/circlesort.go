// Package sort implements various sorting algorithms.
package sort

import "github.com/TheAlgorithms/Go/constraints"

// Circle sorts an array using the circle sort algorithm.
func Circle[T constraints.Ordered](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	for doSort(arr, 0, len(arr)-1) {
	}
	return arr
}

// doSort is the recursive function that implements the circle sort algorithm.
func doSort[T constraints.Ordered](arr []T, left, right int) bool {
	if left == right {
		return false
	}
	swapped := false
	low := left
	high := right

	for low < high {
		if arr[low] > arr[high] {
			arr[low], arr[high] = arr[high], arr[low]
			swapped = true
		}
		low++
		high--
	}

	if low == high && arr[low] > arr[high+1] {
		arr[low], arr[high+1] = arr[high+1], arr[low]
		swapped = true
	}

	mid := left + (right-left)/2
	leftHalf := doSort(arr, left, mid)
	rightHalf := doSort(arr, mid+1, right)

	return swapped || leftHalf || rightHalf
}
