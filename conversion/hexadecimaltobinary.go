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
	"strconv"
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
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "", errors.New("failed to parse hexadecimal string " + hex + ": " + err.Error())
	}

	// Convert the integer to a binary string
	binary := strconv.FormatInt(decimal, 2)
	return binary, nil
}
