// author(s) [jo3zeph](https://github.com/jo3zeph)
// description: Find the median from a set of values
// see median_test.go

package median

import (
	"sort"
)


func median(values []float64) float64 {

	// Sorting the values in order
	sort.Float64s(values)

	var median float64

	// Find the length of values
	l := len(values)

	if l == 0 {
        return 0
    } else if l%2 == 0 {
        median = (values[l/2-1] + values[l/2]) / 2
    } else {
        median = values[l/2]
    }

	return median
}