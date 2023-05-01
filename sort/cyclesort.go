package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
)

// Cycle sort is an in-place, unstable sorting algorithm that is particularly useful
// when sorting arrays containing elements with a small range of values. It is theoretically
// optimal in terms of the total number of writes to the original array.
func Cycle[T constraints.Number](arr []T) []T {
	counter, cycle, len := 0, 0, len(arr)
	// Early return if the array too small
	if len <= 1 {
		return arr
	}

	for cycle = 0; cycle < len-1; cycle++ {
		elem := arr[cycle]
		// Find total smaller elements to right
		pos := cycle
		for counter = cycle + 1; counter < len; counter++ {
			if arr[counter] < elem {
				pos++
			}
		}
		// In case this element is already in correct position, let's skip processing
		if pos == cycle {
			continue
		}
		// In case we have same elements, we want to skip to the end of that list as well, ignoring order
		// This makes the algorithm unstable for composite elements
		for elem == arr[pos] {
			pos++
		}
		// Now let us put the item to it's right position
		arr[pos], elem = elem, arr[pos]

		//We need to rotate the array till we have reached the start of the cycle again
		for pos != cycle {
			pos = cycle
			// Find smaller elements to right again
			for counter = cycle + 1; counter < len; counter++ {
				if arr[counter] < elem {
					pos++
				}
			}
			for elem == arr[pos] {
				pos++
			}
			//We can do this unconditionally, but the check helps prevent redundant writes to the array
			if elem != arr[pos] {
				arr[pos], elem = elem, arr[pos]
			}
		}
	}

	return arr
}
