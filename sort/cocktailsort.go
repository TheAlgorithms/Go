// Implementation of Cocktail sorting
// reference: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

package sort

import "github.com/TheAlgorithms/Go/constraints"

// Cocktail sort is a variation of bubble sort, operating in two directions (beginning to end, end to beginning)
func Cocktail[T constraints.Ordered](arr []T) []T {
	if len(arr) == 0 { // ignore 0 length arrays
		return arr
	}

	swapped := true // true if swapped two or more elements in the last loop
	// if it loops through the array without swapping, the array is sorted

	// start and end indexes, this will be updated excluding already sorted elements
	start := 0
	end := len(arr) - 1

	for swapped {
		swapped = false
		var new_start int
		var new_end int

		for i := start; i < end; i++ { // first loop, from start to end
			if arr[i] > arr[i+1] { // if current and next elements are unordered
				arr[i], arr[i+1] = arr[i+1], arr[i] // swap two elements
				new_end = i
				swapped = true
			}
		}

		end = new_end

		if !swapped { // early exit, skipping the second loop
			break
		}

		swapped = false

		for i := end; i > start; i-- { // second loop, from end to start
			if arr[i] < arr[i-1] { // same process of the first loop, now going 'backwards'
				arr[i], arr[i-1] = arr[i-1], arr[i]
				new_start = i
				swapped = true
			}
		}

		start = new_start
	}

	return arr
}
