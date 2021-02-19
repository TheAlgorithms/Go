package modulararithmetic

import (
	"errors"
	"math"
)

// ErrorIntOverflow For asserting that the values do not overflow in Int64
var ErrorIntOverflow = errors.New("Integer Overflow")

// ModularExponentiation returns base^exponent % mod
func ModularExponentiation(base, exponent, mod int64) int64 {
	if mod == 1 {
		return 0
	}
	_, err := Multiply64BitInt(mod-1, mod-1)

	if err != nil {
		panic(err)
	}

	var result int64 = 1

	base = base % mod

	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		exponent = exponent >> 1
		base = (base * base) % mod
	}
	return result
}

// Multiply64BitInt Checking if the integer multiplication overflows
func Multiply64BitInt(left, right int64) (int64, error) {
	if math.Abs(float64(left)) > float64(math.MaxInt64)/math.Abs(float64(right)) {
		return 0, ErrorIntOverflow
	}
	return left * right, nil
}
