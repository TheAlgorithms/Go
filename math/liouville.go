// liouville.go
// description: Returns λ(n)
// details:
// For any positive integer n, define λ(n) as the sum of the primitive nth roots of unity.
// It has values in {−1, 1} depending on the factorization of n into prime factors:
//   λ(n) = +1 if n is a positive integer with an even number of prime factors.
//   λ(n) = −1 if n is a positive integer with an odd number of prime factors.
// wikipedia: https://en.wikipedia.org/wiki/Liouville_function
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see liouville_test.go

package math

import (
	"errors"

	"github.com/TheAlgorithms/Go/math/prime"
)

var ErrNonZeroArgsOnly error = errors.New("arguments must be greater than zero")

// Lambda is the liouville function
// This function returns λ(n) for given number
func Lambda(n int) (int, error) {
	if n < 0 {
		return 0, ErrPosArgsOnly
	}
	if n == 0 {
		return 0, ErrNonZeroArgsOnly
	}
	if len(prime.Factorize(int64(n)))%2 == 0 {
		return 1, nil
	}
	return -1, nil
}
