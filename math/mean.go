package maths

import (
	"golang.org/x/exp/constraints"
)

type Digit interface {
	constraints.Float | constraints.Integer
}

func Mean[T Digit](values []T) float64 {

	// Validating the empty array
	if len(values) == 0 {
		return 0
	}

	var summation float64 = 0

	// Summation of all values in the slice
	for _, singleValue := range values {
		summation += float64(singleValue)
	}

	// Casting to float64 done for the len()
	var average = summation / float64(len(values))

	// Return of average value
	return average

}
