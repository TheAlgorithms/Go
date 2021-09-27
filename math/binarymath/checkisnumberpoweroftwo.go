// checkisnumberpoweroftwo.go
// description: Is the number a power of two
// details:
// Checks if a number is a power of two- [Power of two](https://en.wikipedia.org/wiki/Power_of_two)
// author(s) [red_byte](https://github.com/i-redbyte)
// see checkisnumberpoweroftwo_test.go

package binarymath

import (
	"math"
)

func IsPowerOfTwoBinaryMethod(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

func IsPowerOfTwoUseCycleAndLeftShift(number uint) bool {
	if number == 0 {
		return false
	}
	for p := uint(1); p > 0; p = p << 1 {
		if number == p {
			return true
		}
	}
	return false
}

func IsPowerOfTwoUseLog(number float64) bool {
	if number == 0 || math.Round(number) == math.MaxInt64 {
		return false
	}
	log := math.Log2(number)
	pow := math.Pow(2, math.Round(log))
	return pow == number
}
