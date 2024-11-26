// Implementation of comb sort algorithm, an improvement of bubble sort
// average time complexity: O(n^2 / 2^p) where p is the number of increments
// worst time complexity: O(n^2)
// space complexity: O(1)
// Reference: https://www.geeksforgeeks.org/comb-sort/

package sort

import "github.com/TheAlgorithms/Go/constraints"

func getNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}

// Comb is a simple sorting algorithm which is an improvement of the bubble sorting algorithm.
func Comb[T constraints.Ordered](data []T) []T {
	n := len(data)
	gap := n
	swapped := true

	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		for i := 0; i < n-gap; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
				swapped = true
			}
		}
	}
	return data
}
