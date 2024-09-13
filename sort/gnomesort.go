// gnomesort.go
// description: implementation of the Gnome sort algorithm in Go
// details:
// Gnome sort is an algorithm similar to insertion sort that uses backward swaps to reorder. [Wikipedia] (https://en.wikipedia.org/wiki/Gnome_sort)
// author [Ibrahim](https://github.com/ibrahimelnemr)
// see sort_test.go for testing of this algorithm, test function TestGnome

package sort

import "github.com/TheAlgorithms/Go/constraints"

// GnomeSort operates similarly to insertion sort but uses backward swaps and has no nested loops.
func Gnome[T constraints.Ordered](arr []T) []T {
	i := 0
	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
	return arr
}
