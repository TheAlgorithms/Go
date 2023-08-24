package sqrt

import (
	"unsafe"
)

// Assumes that number always positive
// You don't want to deal with complex numbers
func Q_rsqrt(number float32) float32 {
	var i int32
	var y float32
	const threehalfs float32 = 1.5

	var x2 float32 = number * float32(0.5)
	y = number
	i = *(*int32)(unsafe.Pointer(&y)) // evil floating point bit level hacking
	i = 0x5f3759df - (i >> 1)         // what the fuck?
	y = *(*float32)(unsafe.Pointer(&i))

	y = y * (threehalfs - (x2 * y * y)) // 1st iteration
	y = y * (threehalfs - (x2 * y * y)) // 2nd iteration, this can be removed
	return y
}
