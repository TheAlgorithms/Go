/**
* Problem 6 - Sum square difference
* @see {@link https://projecteuler.net/problem=6}
*
* The sum of the squares of the first ten natural numbers is,
* 1^2 + 2^2 + ... + 10^2 = 385
*
* The square of the sum of the first ten natural numbers is,
* (1 + 2 + ... + 10)^2 = 55^2 = 3025
*
* Hence the difference between the sum of the squares of the first ten natural numbers
* and the square of the sum is 3025 − 385 = 2640.
*
* Find the difference between the sum of the squares of the first one hundred natural numbers
* and the square of the sum.
*
* @author ddaniel27
 */
package problem6

func Problem6(n uint) uint {
	sumOfSquares := uint(0)
	squareOfSum := uint(0)

	for i := uint(1); i <= n; i++ {
		sumOfSquares += i * i
		squareOfSum += i
	}

	squareOfSum *= squareOfSum

	return squareOfSum - sumOfSquares
}
