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
