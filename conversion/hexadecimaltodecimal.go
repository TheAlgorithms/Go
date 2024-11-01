/*
Author: mapcrafter2048
GitHub: https://github.com/mapcrafter2048
*/

// This algorithm will convert any Hexadecimal number(0-9, A-F, a-f) to Decimal number(0-9).
// https://en.wikipedia.org/wiki/Hexadecimal
// https://en.wikipedia.org/wiki/Decimal
// Function receives a Hexadecimal Number as string and returns the Decimal number as integer.
// Supported Hexadecimal number range is 0 to 7FFFFFFFFFFFFFFF.

package conversion

import (
	"fmt"
	"regexp"
	"strings"
)

var isValidHexadecimal = regexp.MustCompile("^[0-9A-Fa-f]+$").MatchString

// hexToDecimal converts a hexadecimal string to a decimal integer.
func hexToDecimal(hexStr string) (int64, error) {

	hexStr = strings.TrimSpace(hexStr)

	if len(hexStr) == 0 {
		return 0, fmt.Errorf("input string is empty")
	}

	// Check if the string has a valid hexadecimal prefix
	if len(hexStr) > 2 && (hexStr[:2] == "0x" || hexStr[:2] == "0X") {
		hexStr = hexStr[2:]
	}

	// Validate the hexadecimal string
	if !isValidHexadecimal(hexStr) {
		return 0, fmt.Errorf("invalid hexadecimal string")
	}

	var decimalValue int64
	for _, char := range hexStr {
		var digit int64
		if char >= '0' && char <= '9' {
			digit = int64(char - '0')
		} else if char >= 'A' && char <= 'F' {
			digit = int64(char - 'A' + 10)
		} else if char >= 'a' && char <= 'f' {
			digit = int64(char - 'a' + 10)
		} else {
			return 0, fmt.Errorf("invalid character in hexadecimal string: %c", char)
		}
		decimalValue = decimalValue*16 + digit
	}

	return decimalValue, nil
}
