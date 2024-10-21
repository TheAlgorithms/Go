/*
Author: Motasim
GitHub: https://github.com/motasimmakki
Date: 19-Oct-2021
*/

// This algorithm will convert any Binary number(0 or 1) to Decimal number(+ve number).
// https://en.wikipedia.org/wiki/Binary_number
// https://en.wikipedia.org/wiki/Decimal
// Function receives a Binary Number as string and returns the Decimal number as integer.
// Supported Binary number range is 0 to 2^(31-1).
// time complexity: O(n)
// space complexity: O(1)

package conversion

// Importing necessary package.
import (
	"errors"
	"regexp"
)

var isValid = regexp.MustCompile("^[0-1]{1,}$").MatchString

// BinaryToDecimal() function that will take Binary number as string,
// and return its Decimal equivalent as an integer.
func BinaryToDecimal(binary string) (int, error) {
	if !isValid(binary) {
		return -1, errors.New("not a valid binary string")
	}
	if len(binary) > 32 {
		return -1, errors.New("binary number must be in range 0 to 2^(31-1)")
	}
	var result, base int = 0, 1
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			result += base
		}
		base *= 2
	}
	return result, nil
}
