// This algorithm will convert a standard roman number to an integer
// https://en.wikipedia.org/wiki/Roman_numerals
// Function receives a string as a roman number and outputs an integer
// Maximum output will be 3999
// Only standard form is supported
// time complexity: O(n)
// space complexity: O(1)

package conversion

import (
	"errors"
	"strings"
)

// numeral describes the value and symbol of a single roman numeral
type numeral struct {
	val int
	sym string
}

// lookup array for numeral values sorted by largest to smallest
var nums = []numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// RomanToInt converts a roman numeral string to an integer. Roman numerals for numbers
// outside the range 1 to 3,999 will return an error. Nil or empty string return 0
// with no error thrown.
func RomanToInt(input string) (int, error) {
	if input == "" {
		return 0, nil
	}
	var output int
	for _, n := range nums {
		for strings.HasPrefix(input, n.sym) {
			output += n.val
			input = input[len(n.sym):]
		}
	}
	// if we are still left with input string values then the
	// input was invalid and an error is returned.
	if len(input) > 0 {
		return 0, errors.New("invalid roman numeral")
	}
	return output, nil
}
