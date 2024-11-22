/**
* Problem 5 - Smallest multiple
* @see {@link https://projecteuler.net/problem=5}
*
* 2520 is the smallest number that can be divided by
* each of the numbers from 1 to 10 without any remainder.
* What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*
* @author ddaniel27
 */
package problem5

func Problem5(limit uint) uint {
	n := limit * limit

	for {
		if isDivisible(n, limit) {
			return n
		}

		n++
	}
}

func isDivisible(n, limit uint) bool {
	for i := uint(1); i <= limit; i++ {
		if n%i != 0 {
			return false
		}
	}

	return true
}
