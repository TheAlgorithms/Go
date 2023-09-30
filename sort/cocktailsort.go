// Implementation of Cocktail sorting
// reference: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

// Cocktail sort is a variation of bubble sort, operating in two directions (beginning to end, end to beginning)
func Cocktail[T constraints.Integer](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	swapped := false

	for ok := true; ok; ok = swapped {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false

		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
