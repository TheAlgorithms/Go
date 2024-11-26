// inverse.go
// description: Implementation of Modular Inverse Algorithm
// details:
// A simple implementation of Modular Inverse - [Modular Inverse wiki](https://en.wikipedia.org/wiki/Modular_multiplicative_inverse)
// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)
// author(s) [Taj](https://github.com/tjgurwara99)
// see inverse_test.go

package modular

import (
	"errors"

	"github.com/TheAlgorithms/Go/math/gcd"
)

var ErrorInverse = errors.New("no Modular Inverse exists")

// Inverse Modular function
func Inverse(a, m int64) (int64, error) {
	gcd, x, _ := gcd.Extended(a, m)
	if gcd != 1 || m == 0 {
		return 0, ErrorInverse
	}

	return ((m + (x % m)) % m), nil // this is necessary because of Go's use of architecture specific instruction for the % operator.
}
