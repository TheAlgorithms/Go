package lcm

import (
	"math"

	"github.com/TheAlgorithms/Go/math/gcd"
)

// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd.Iterative(a, b)))
}
