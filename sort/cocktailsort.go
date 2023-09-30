// Implementation of Cocktail sorting
// reference: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

func Cocktail[T constraints.Integer](arr []T) []T {
	if len(arr) == 0 { // guard clause for 0 length arrays
		return arr
	}

	swapped := false

	for ok := true; ok; ok = swapped {
		for i := 0; i < len(arr)-1; i++ { // loop from array's beginning to end
			if arr[i] > arr[i+1] { // swap unsorted elements
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			// if swapped == false, first loop didn't find any unsorted elements.
			// array is sorted
			break
		}

		swapped = false

		for i := len(arr) - 1; i > 0; i-- { // loop from array's end to beginning
			if arr[i] < arr[i-1] { // swap unsorted elements
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
