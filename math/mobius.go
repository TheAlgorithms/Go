// mobius.go
// description: Returns μ(n)
// details:
// For any positive integer n, define μ(n) as the sum of the primitive nth roots of unity.
// It has values in {−1, 0, 1} depending on the factorization of n into prime factors:
//   μ(n) = +1 if n is a square-free positive integer with an even number of prime factors.
//   μ(n) = −1 if n is a square-free positive integer with an odd number of prime factors.
//   μ(n) = 0 if n has a squared prime factor.
// wikipedia: https://en.wikipedia.org/wiki/M%C3%B6bius_function
// time complexity: O(n)
// space complexity: O(1)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see mobius_test.go

package math

import (
	"github.com/TheAlgorithms/Go/math/prime"
)

// Mu is the Mobius function
// This function returns μ(n) for given number
func Mu(n int) int {
	if n <= 1 {
		return 1
	}
	var primeFactorCount int
	for i := 1; i <= n; i++ {
		if n%i == 0 && prime.OptimizedTrialDivision(int64(i)) {
			if n%(i*i) == 0 {
				return 0
			}
			primeFactorCount += 1
		}
	}
	if primeFactorCount%2 == 0 {
		return 1
	}
	return -1
}
