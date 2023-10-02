// Implementation of Cocktail sorting
// reference: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

// Cocktail sort is a variation of bubble sort, operating in two directions (beginning to end, end to beginning)
func Cocktail[T constraints.Integer](arr []T) []T {
	if len(arr) == 0 { // ignore 0 length arrays
		return arr
	}

	swapped := false // true if swapped two or more elements in the last loop
	// if it loops through the array without swapping, the array is sorted

	for ok := true; ok; ok = swapped { // exits when swapped == false

		for i := 0; i < len(arr)-1; i++ { // first loop, from array's beginning to end
			if arr[i] > arr[i+1] { // if current and next elements are unordered
				arr[i], arr[i+1] = arr[i+1], arr[i] // swap two elements
				swapped = true
			}
		}

		if !swapped { // early exit, skipping the second loop
			break
		}

		swapped = false

		for i := len(arr) - 1; i > 0; i-- { // second loop, from array's end to beginning
			if arr[i] < arr[i-1] { // same process of the first loop, now going 'backwards'
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
