// author(s) [jo3zeph](https://github.com/jo3zeph)
// description: Find the median from a set of values
// time complexity: O(n log n)
// space complexity: O(1)
// see median_test.go

package math

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/sort"
)

func Median[T constraints.Number](values []T) float64 {

	sort.Bubble(values)

	l := len(values)

	switch {
	case l == 0:
		return 0

	case l%2 == 0:
		return float64((values[l/2-1] + values[l/2]) / 2)

	default:
		return float64(values[l/2])
	}
}
