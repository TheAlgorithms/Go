package sort

import "github.com/TheAlgorithms/Go/constraints"

// Pancake sorts a slice using flip operations,
// where flip refers to the idea of reversing the
// slice from index `0` to `i`.
func Pancake[T constraints.Ordered](arr []T) []T {
	// early return if the array too small
	if len(arr) <= 1 {
		return arr
	}

	// start from the end of the array
	for i := len(arr) - 1; i > 0; i-- {
		// find the index of the maximum element in arr
		max := 0
		for j := 1; j <= i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}

		// if the maximum element is not at the end of the array
		if max != i {
			// flip the maximum element to the beginning of the array
			arr = flip(arr, max)

			// flip the maximum element to the end of the array by flipping the whole array
			arr = flip(arr, i)
		}
	}

	return arr
}

// flip reverses the input slice from `0` to `i`.
func flip[T constraints.Ordered](arr []T, i int) []T {
	for j := 0; j < i; j++ {
		arr[j], arr[i] = arr[i], arr[j]
		i--
	}
	return arr
}
