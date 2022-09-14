// binomialcoefficient.go
// description: Returns C(n, k)
// details:
// a binomial coefficient C(n,k) gives number ways
// in which k objects can be chosen from n objects.
// wikipedia: https://en.wikipedia.org/wiki/Binomial_coefficient
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see binomialcoefficient_test.go

package math

import (
	"errors"
)

// C is Binomial Coefficient function
// This function returns C(n, k) for given n and k
func Combinations(n int, k int) (int, error) {
	var ErrPosArgsOnly error = errors.New("arguments must be positive")
	if n < 0 || k < 0 {
		return -1, ErrPosArgsOnly
	}
	if k > (n - k) {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res *= (n - i)
		res /= (i + 1)
	}
	return res, nil
}
