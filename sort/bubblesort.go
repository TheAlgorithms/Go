// Implementation of basic bubble sort algorithm
// Reference: https://en.wikipedia.org/wiki/Bubble_sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

// Bubble is a simple generic definition of Bubble sort algorithm.
func Bubble[T constraints.Ordered](arr []T) []T {
	swapped := true
	sortedIdx := 1
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-sortedIdx; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
