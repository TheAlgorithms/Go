package search

import (
	"fmt"
)

// Package search provides functions for searching and root-finding algorithms.
//
// bisection searches for the root of the function f in the interval [a, b].
// The function f must be continuous and f(a) and f(b) must have different signs.
// This method uses the bisection algorithm to iteratively narrow down the interval
// until the root is found within the specified tolerance or the maximum number of
// iterations is reached.
//
// Parameters:
//   - f: A function that takes a float64 and returns a float64. This is the function
//     for which we are trying to find the root.
//   - a: The start of the interval.
//   - b: The end of the interval.
//   - tol: The tolerance for the root. The algorithm stops when the interval width
//     is less than this value.
//   - maxIter: The maximum number of iterations to perform.
//
// Returns:
//   - float64: The approximate root of the function.
//   - error: An error is returned if f(a) and f(b) do not have different signs or
//     if the maximum number of iterations is reached without finding the root.
//
// Complexity:
//   - Worst: O(log((b-a)/tol))
//   - Average: O(log((b-a)/tol))
//   - Best: O(1)
//
// Example usage:
//
//	root, err := bisection(func(x float64) float64 { return math.Pow(x, 3) - x - 2 }, 1, 2, 1e-6, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(root) // Output: 1.521379
//
// Reference:
//   - https://en.wikipedia.org/wiki/Bisection_method
func bisection(f func(float64) float64, a, b, tol float64, maxIter int) (float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, fmt.Errorf("f(a) and f(b) must have different signs")
	}

	var c float64
	for i := 0; i < maxIter; i++ {
		c = (a + b) / 2
		if f(c) == 0 || (b-a)/2 < tol {
			return c, nil
		}
		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return c, fmt.Errorf("maximum number of iterations reached")
}
