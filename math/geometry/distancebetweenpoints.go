// distancebetweenpoints.go
// Find distance between two points
// Details:
// Given coordinates of 2 points on a cartesian plane, find the distance between them rounded up to nearest integer.
// Link: https://practice.geeksforgeeks.org/problems/distance-between-2-points3200/1?page=1&category[]=geometric

// author(s) [Chetan Patil](https://github.com/Chetan07j)

// Package geometry contains geometric algorithms
package geometry

import (
	"errors"
	"math"
)

// Point defines a point with x and y coordinates.
type CartesianPoint []float64

var ErrDimMismatch = errors.New("mismatched dimensions")

// This function calculates the distance based on Cartesian Plane Distance Formula
// Formula: √((x1-x2) * (x1-x2) + (y1-y2) * (y1-y2))
// Takes x and y coordinates of two point
// Returns distance in floating number
func FindDistanceBetweenTwoPoints(p1 CartesianPoint, p2 CartesianPoint) (float64, error) {
	// Find length of point 1
	p1Length := len(p1)

	// Verify length of point 1 and 2
	// and throw error if length do not match
	if len(p2) != p1Length {
		return -1, ErrDimMismatch
	}

	var total float64 = 0

	// Iterate over point1
	// find x2 - x1 and y2 - y1 as per formula
	// here we are doing x1 - x1 and y1 - y2
	// math.Abs is applied so that we do not get negative value
	for i, x_i := range p1 {
		// Find difference
		diff := math.Abs(x_i - p2[i])

		// Add diff * diff to total value
		total += diff * diff
	}

	// Last step of formula
	// return square root of total
	return math.Sqrt(total), nil
}
