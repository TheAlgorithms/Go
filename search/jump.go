// jump.go
// description: Implementation of jump search
// details:
// A search algorithm for ordered list that jump through the list to narrow down the range
// before performing a linear search
// reference: https://en.wikipedia.org/wiki/Jump_search
// see jump_test.go for a test implementation, test function TestJump
// time complexity: O(sqrt(n))
// space complexity: O(1)

package search

import "math"

// Jump search works by jumping multiple steps ahead in sorted list until it find an item larger than target,
// then create a sublist of item from the last searched item up to the current item and perform a linear search.
func Jump(array []int, target int) (int, error) {
	n := len(array)
	if n == 0 {
		return -1, ErrNotFound
	}

	// the optimal value of step is square root of the length of list
	step := int(math.Round(math.Sqrt(float64(n))))

	prev := 0    // previous index
	curr := step // current index
	for array[curr-1] < target {
		prev = curr
		if prev >= len(array) {
			return -1, ErrNotFound
		}

		curr += step

		// prevent jumping over list range
		if curr > n {
			curr = n
		}
	}

	// perform linear search from index prev to index curr
	for array[prev] < target {
		prev++

		// if reach end of range, indicate target not found
		if prev == curr {
			return -1, ErrNotFound
		}
	}
	if array[prev] == target {
		return prev, nil
	}

	return -1, ErrNotFound

}
