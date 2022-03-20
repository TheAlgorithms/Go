// author(s) [red_byte](https://github.com/i-redbyte)
// see sin_test.go

package math

import "math"

// Sin returns the sine of the radian argument x. [See more](https://en.wikipedia.org/wiki/Sine_and_cosine)
func Sin(x float64) float64 {
	return Cos((math.Pi / 2) - x)
}
