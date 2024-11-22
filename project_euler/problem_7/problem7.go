/**
* Problem 7 - 10001st prime
* @see {@link https://projecteuler.net/problem=7}
*
* By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
* we can see that the 6th prime is 13.
*
* What is the 10 001st prime number?
*
* @author ddaniel27
 */
package problem7

import "github.com/TheAlgorithms/Go/math/prime"

func Problem7(n uint) int64 {
	count, i := uint(0), int64(1)

	for count < n {
		i++
		if prime.OptimizedTrialDivision(i) {
			count++
		}
	}

	return i
}
