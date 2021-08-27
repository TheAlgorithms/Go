package modular

import (
	"errors"

	"github.com/TheAlgorithms/Go/math/gcd"
)

// ErrorIntOverflow For asserting that the values do not overflow in Int64
var ErrorInverse = errors.New("No Modular Inverse Exists")

// Inverse Modular function
func Inverse(a, b int64) (int64, error) {
	gcd, x, _ := gcd.Extended(a, b)
	if gcd != 1 {
		return 0, ErrorInverse
	}

	return ((b + (x % b)) % b), nil // this is necessary because of Go's use of architecture specific instruction for the % operator.
}
