// sqrt.go
// description: Square root calculation
// details:
// Calculating the square root using binary operations and a magic number 0x5f3759df [See more](https://en.wikipedia.org/wiki/Fast_inverse_square_root)
// author(s) [red_byte](https://github.com/i-redbyte)
// see sqrt_test.go

package binary

import (
	"math"
)

const threeHalves = 1.5

func Sqrt(number float32) float64 {
	var halfHumber, y float32
	halfHumber = number * 0.5
	z := math.Float32bits(number) // floating point bit level hacking
	z = 0x5f3759df - (z >> 1)     // Newton's approximation
	y = math.Float32frombits(z)
	for i := 0; i < 3; i++ {
		y = y * (threeHalves - (halfHumber * y * y))
	}
	return math.RoundToEven(float64(1 / y))
}
