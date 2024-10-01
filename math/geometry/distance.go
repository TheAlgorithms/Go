// distance.go
// Find Euclidean distance between two points
// time complexity: O(n) where n is the number of dimensions
// space complexity: O(1)
// author(s) [Chetan Patil](https://github.com/Chetan07j)

// Package geometry contains geometric algorithms
package geometry

import (
	"errors"
	"math"
)

// EuclideanPoint defines a point with x and y coordinates.
type EuclideanPoint []float64

var ErrDimMismatch = errors.New("mismatched dimensions")

// EuclideanDistance returns the Euclidean distance between points in
// any `n` dimensional Euclidean space.
func EuclideanDistance(p1 EuclideanPoint, p2 EuclideanPoint) (float64, error) {
	n := len(p1)

	if len(p2) != n {
		return -1, ErrDimMismatch
	}

	var total float64 = 0

	for i, x_i := range p1 {
		// using Abs since the value could be negative but we require the magnitude
		diff := math.Abs(x_i - p2[i])
		total += diff * diff
	}

	return math.Sqrt(total), nil
}
