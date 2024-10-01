// Calculating the inverse square root
// time complexity: O(1)
// space complexity: O(1)
// [See more](https://en.wikipedia.org/wiki/Fast_inverse_square_root)

package binary

import (
	"math"
)

// FastInverseSqrt assumes that argument is always positive,
// and it does not deal with negative numbers.
// The "magic" number 0x5f3759df is hex for 1597463007 in decimals.
// The math.Float32bits is alias to *(*uint32)(unsafe.Pointer(&f))
// and math.Float32frombits is to *(*float32)(unsafe.Pointer(&b)).
func FastInverseSqrt(number float32) float32 {
	var i uint32
	var y, x2 float32
	const threehalfs float32 = 1.5

	x2 = number * float32(0.5)
	y = number
	i = math.Float32bits(y)   // evil floating point bit level hacking
	i = 0x5f3759df - (i >> 1) // magic number and bitshift hacking
	y = math.Float32frombits(i)

	y = y * (threehalfs - (x2 * y * y)) // 1st iteration of Newton's method
	y = y * (threehalfs - (x2 * y * y)) // 2nd iteration, this can be removed
	return y
}
