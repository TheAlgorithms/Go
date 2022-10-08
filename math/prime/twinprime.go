// twinprime.go
// description: Returns Twin Prime of n
// details:
// For any integer n, twin prime is (n + 2)
// if and only if both n and (n + 2) both are prime
// wikipedia: https://en.wikipedia.org/wiki/Twin_prime
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see twinprime_test.go

package math

import (
	"github.com/TheAlgorithms/Go/math/prime"
)

// This function returns twin prime for given number
// returns (n + 2) if both n and (n + 2) are prime
// -1 otherwise
func TwinPrime(n int) int {
	if prime.OptimizedTrialDivision(int64(n)) && prime.OptimizedTrialDivision(int64(n+2)) {
		return n + 2
	}
	return -1
}
