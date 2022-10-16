package max

import "github.com/TheAlgorithms/Go/constraints"

// Int is a function which returns the maximum of all the integers provided as arguments.
func Int[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
