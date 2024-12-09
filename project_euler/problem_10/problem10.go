/**
* Problem 10 - Summation of primes
* @see {@link https://projecteuler.net/problem=10}
*
* The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
* Find the sum of all the primes below two million.
*
* @author ddaniel27
 */
package problem10

import "github.com/TheAlgorithms/Go/math/prime"

func Problem10(n int) uint {
	sum := uint(0)
	sieve := prime.SieveEratosthenes(n)

	for _, v := range sieve {
		sum += uint(v)
	}

	return sum
}
