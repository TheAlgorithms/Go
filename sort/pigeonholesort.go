package sort

import (
	"github.com/TheAlgorithms/Go/math/max"
	"github.com/TheAlgorithms/Go/math/min"
)

// Pigeonhole algorithm's working at wikipedia.
// https://en.wikipedia.org/wiki/Pigeonhole_sort
func Pigeonhole(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := max.Int(arr...)
	min := min.Int(arr...)

	size := max - min + 1

	holes := make([]int, size)

	for _, element := range arr {
		holes[element-min]++
	}

	i := 0

	for j := 0; j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			arr[i] = j + min
			i++
		}
	}

	return arr
}
