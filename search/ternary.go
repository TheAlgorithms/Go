package search

import (
	"fmt"
	"math"
)

// TernaryMax is a function to search for maximum value of a uni-modal function `f`
// in the interval [a, b]. a and b should be finit numbers
func TernaryMax(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval boundaries should be finite numbers")
	}
	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}
	left := (2*a + b) / 3
	right := (a + 2*b) / 3
	if f(left) < f(right) {
		return TernaryMax(left, b, epsilon, f)
	}
	return TernaryMax(a, right, epsilon, f)
}

// TernaryMin is a function to search for minimum value of a uni-modal function `f`
// in the interval [a, b]. a and b should be finit numbers.
func TernaryMin(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval boundaries should be finite numbers")
	}
	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}
	left := (2*a + b) / 3
	right := (a + 2*b) / 3
	if f(left) > f(right) {
		return TernaryMin(left, b, epsilon, f)
	}
	return TernaryMin(a, right, epsilon, f)
}
