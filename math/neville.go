// Package math provides common mathematical algorithms and functions.
package math

import "errors"

// Neville implements Neville's algorithm for polynomial interpolation.
// It computes the value of the unique polynomial of degree at most n that
// passes through n+1 given points (xCoords, yCoords) at a specific value `target`.
// For more details, see: https://en.wikipedia.org/wiki/Neville%27s_algorithm
//
// Author: [Mitrajit Ghorui](https://github.com/KeyKyrios)
func Neville(xCoords, yCoords []float64, target float64) (float64, error) {
	if len(xCoords) != len(yCoords) {
		return 0, errors.New("x and y coordinate slices must have the same length")
	}
	if len(xCoords) == 0 {
		return 0, errors.New("input slices cannot be empty")
	}

	seenX := make(map[float64]bool)
	for _, x := range xCoords {
		if seenX[x] {
			return 0, errors.New("input x-coordinates must be unique")
		}
		seenX[x] = true
	}

	n := len(xCoords)
	p := make([]float64, n)
	copy(p, yCoords)

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			p[i] = ((target-xCoords[i+k])*p[i] + (xCoords[i]-target)*p[i+1]) / (xCoords[i] - xCoords[i+k])
		}
	}

	return p[0], nil
}