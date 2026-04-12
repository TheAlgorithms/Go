package search

import (
	"math/bits"
	"sort"
)

// SelectK returns the k-th largest element in array.
//
// Time complexity is expected O(n) thanks to quickselect partitioning.
// A depth limit is applied and falls back to sorting the narrowed range,
// which guarantees O(n log n) worst-case behavior.
//
// The function mutates the input slice in-place.
func SelectK(array []int, k int) (int, error) {
	n := len(array)
	if n == 0 || k < 1 || k > n {
		return -1, ErrNotFound
	}

	// k-th largest -> index in zero-based ascending order.
	idx := n - k
	return selectK(array, idx), nil
}

// selectK returns the element that would appear at idx in sorted order.
func selectK(array []int, idx int) int {
	l, r := 0, len(array)
	depthLimit := 2 * bits.Len(uint(len(array)))

	for r-l > 1 {
		if depthLimit == 0 {
			sort.Ints(array[l:r])
			return array[idx]
		}
		depthLimit--

		leftPivot, rightPivot := partition3(array, l, r)
		switch {
		case idx < leftPivot:
			r = leftPivot
		case idx >= rightPivot:
			l = rightPivot
		default:
			return array[idx]
		}
	}

	return array[l]
}

// partition3 applies a Dutch National Flag partition around a pivot and
// returns [leftPivot, rightPivot), the range that equals the pivot.
func partition3(array []int, l, r int) (leftPivot, rightPivot int) {
	pivotIdx := medianOfThreeIndex(array, l, l+(r-l)/2, r-1)
	pivot := array[pivotIdx]
	array[l], array[pivotIdx] = array[pivotIdx], array[l]

	lt, i, gt := l, l+1, r
	for i < gt {
		switch {
		case array[i] < pivot:
			lt++
			array[i], array[lt] = array[lt], array[i]
			i++
		case array[i] > pivot:
			gt--
			array[i], array[gt] = array[gt], array[i]
		default:
			i++
		}
	}

	array[l], array[lt] = array[lt], array[l]
	return lt, gt
}

func medianOfThreeIndex(array []int, a, b, c int) int {
	if array[a] < array[b] {
		if array[b] < array[c] {
			return b
		}
		if array[a] < array[c] {
			return c
		}
		return a
	}

	if array[a] < array[c] {
		return a
	}
	if array[b] < array[c] {
		return c
	}
	return b
}
