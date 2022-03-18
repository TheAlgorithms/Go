// Package trigan for trigonometric functions and algorithms
package trigan

import "math"

// Cos  returns the cosine of the radian argument x. [See more](https://en.wikipedia.org/wiki/Sine_and_cosine)
// author(s) [red_byte](https://github.com/i-redbyte)
// see sqrt_test.go

func Cos(x float64) float64 {
	tp := 1. / (2. * math.Pi)
	x *= tp
	x -= .25 + math.Floor(x+.25)
	x *= 16. * (math.Abs(x) - .5)
	x += .225 * x * (math.Abs(x) - 1.) //Extra precision
	return x
}
