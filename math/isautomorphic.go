// isautomorphic.go
// description: Checks whether a whole number integer is Automorphic or not. If number < 0 then returns false.
// details:
// In mathematics, a number n is said to be a Automorphic number if the square of n ends in the same digits as n itself.
// ref: (https://en.wikipedia.org/wiki/Automorphic_number)
// time complexity: O(log10(N))
// space complexity: O(1)
// author: [SilverDragonOfR](https://github.com/SilverDragonOfR)

package math

import (
	"github.com/TheAlgorithms/Go/constraints"
)

func IsAutomorphic[T constraints.Integer](n T) bool {
	// handling the negetive number case
	if n < 0 {
		return false
	}

	n_sq := n * n
	for n > 0 {
		if (n % 10) != (n_sq % 10) {
			return false
		}
		n /= 10
		n_sq /= 10
	}

	return true
}
