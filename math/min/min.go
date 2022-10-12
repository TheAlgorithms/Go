package min

import "github.com/TheAlgorithms/Go/constraints"

// Int is a function which returns the minimum of all the integers provided as arguments.
func Int[T constraints.Integer](values ...T) T {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
