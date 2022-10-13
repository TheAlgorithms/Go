// distancebetweenpoints.go
// Find distance between two points
// Details:
// Given coordinates of 2 points on a cartesian plane, find the distance between them rounded up to nearest integer.
// Link: https://practice.geeksforgeeks.org/problems/distance-between-2-points3200/1?page=1&category[]=geometric

// author(s) [Chetan Patil](https://github.com/Chetan07j)

// Package geometry contains geometric algorithms
package geometry

import (
	"math"
)

// This function calculates the distance based on Cartesian Plane Distance Formula
// Formula: √((x1-x2) * (x1-x2) + (y1-y2) * (y1-y2))
// Takes x and y coordinates of two point
// Returns distnce in floating number
func FindDistanceBetweenTwoPoints(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	// Find difference between x-coordinates
	xDiff := x2 - x1

	// Find difference between y-coordinates
	yDiff := y2 - y1

	// Apply formula
	distance := (xDiff * xDiff) + (yDiff * yDiff)

	// Return square root
	return math.Sqrt(distance)
}
