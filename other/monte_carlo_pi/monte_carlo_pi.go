// Monte Carlo Pi
// [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method) to calculate Pi
//
// The basic idea of Monte Carlo is:
//  1. Define a domain of possible inputs
//  2. Generate inputs randomly from a probability distribution over the domain
//  3. Perform a deterministic computation on the inputs
//  4. Aggregate the results
//
// Thus the algorithm to calculate Pi is:
//
// 1. Given a circle inscribed in a square with side length 1.
// 2. randomly draw n points (x, y) | x,y in [0.0,1.0)
// 3. Count the amount of points inside the circle
// 4. pi = 4 * (points inside the circle) / n
//
// A point (x, y) is inside the circle if:
//
// x*x + y*y <= 1
//

package main

import (
	"math"
	"math/rand"
)

// MonteCarloPi calculates Pi using the Monte Carlo method
func MonteCarloPi(iterations uint) float64 {
	var inside, total uint

	for total = uint(0); total < iterations; total++ {
		if in := math.Hypot(rand.Float64(), rand.Float64()) <= 1; in {
			inside++
		}
	}

	return float64(4) * (float64(inside) / float64(total))
}
