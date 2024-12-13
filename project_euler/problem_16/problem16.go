/**
* Problem 16 - Power digit sum
* @see {@link https://projecteuler.net/problem=16}
*
* 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
*
* What is the sum of the digits of the number 2^1000?
*
* @author ddaniel27
 */
package problem16

import (
	"math/big"
)

func Problem16(exponent int64) int64 {
	var result big.Int

	bigTwo := big.NewInt(2)
	bigExponent := big.NewInt(exponent)

	result.Exp(bigTwo, bigExponent, nil)

	resultStr := result.String()

	var sum int64
	for _, digit := range resultStr {
		sum += int64(digit - '0')
	}

	return sum
}
