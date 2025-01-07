/*
Author: mapcrafter2048
GitHub: https://github.com/mapcrafter2048
*/

// This algorithm will convert any Hexadecimal number(0-9, A-F, a-f) to Binary number(0 or 1).
// https://en.wikipedia.org/wiki/Hexadecimal
// https://en.wikipedia.org/wiki/Binary_number
// Function receives a Hexadecimal Number as string and returns the Binary number as string.
// Supported Hexadecimal number range is 0 to 7FFFFFFFFFFFFFFF.

package conversion

import (
	"errors"
	"regexp"
	"strings"
)

var isValidHex = regexp.MustCompile("^[0-9A-Fa-f]+$").MatchString

// hexToBinary() function that will take Hexadecimal number as string,
// and return its Binary equivalent as a string.
func hexToBinary(hex string) (string, error) {
	// Trim any leading or trailing whitespace
	hex = strings.TrimSpace(hex)

	// Check if the hexadecimal string is empty
	if hex == "" {
		return "", errors.New("input string is empty")
	}

	// Check if the hexadecimal string is valid
	if !isValidHex(hex) {
		return "", errors.New("invalid hexadecimal string: " + hex)
	}

	// Parse the hexadecimal string to an integer
	var decimal int64
	for i := 0; i < len(hex); i++ {
		char := hex[i]
		var value int64
		if char >= '0' && char <= '9' {
			value = int64(char - '0')
		} else if char >= 'A' && char <= 'F' {
			value = int64(char - 'A' + 10)
		} else if char >= 'a' && char <= 'f' {
			value = int64(char - 'a' + 10)
		} else {
			return "", errors.New("invalid character in hexadecimal string: " + string(char))
		}
		decimal = decimal*16 + value
	}

	// Convert the integer to a binary string without using predefined functions
	var binaryBuilder strings.Builder
	if decimal == 0 {
		binaryBuilder.WriteString("0")
	} else {
		for decimal > 0 {
			bit := decimal % 2
			if bit == 0 {
				binaryBuilder.WriteString("0")
			} else {
				binaryBuilder.WriteString("1")
			}
			decimal = decimal / 2
		}
	}

	// Reverse the binary string since the bits are added in reverse order
	binaryRunes := []rune(binaryBuilder.String())
	for i, j := 0, len(binaryRunes)-1; i < j; i, j = i+1, j-1 {
		binaryRunes[i], binaryRunes[j] = binaryRunes[j], binaryRunes[i]
	}

	return string(binaryRunes), nil
}
