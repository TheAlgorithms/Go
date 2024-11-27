/**
* Problem 3 - Largest prime factor
* @see {@link https://projecteuler.net/problem=3}
*
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 600851475143 ?
*
* @author ddaniel27
 */
package problem3

func Problem3(n uint) uint {
	i := uint(2)

	for n > 1 {
		if n%i == 0 {
			n /= i
		} else {
			i++
		}
	}

	return i
}
