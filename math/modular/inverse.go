// inverse.go
// description: Implementation of Modular Inverse Algorithm
// details:
// A simple implementation of Modular Inverse - [Modular Inverse wiki](https://en.wikipedia.org/wiki/Modular_multiplicative_inverse)
// author(s) [Taj](https://github.com/tjgurwara99)
// see inverse_test.go

package modular

import (
	"errors"

	"github.com/TheAlgorithms/Go/math/gcd"
)

// ErrorIntOverflow For asserting that the values do not overflow in Int64
var ErrorInverse = errors.New("no Modular Inverse exists")

// Inverse Modular function
func Inverse(a, b int64) (int64, error) {
	gcd, x, _ := gcd.Extended(a, b)
	if gcd != 1 {
		return 0, ErrorInverse
	}

	return ((b + (x % b)) % b), nil // this is necessary because of Go's use of architecture specific instruction for the % operator.
}
