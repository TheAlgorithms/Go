// a1z26.go
// description: Implementation of the a1z26 cipher
// details:
// The A1Z26 cipher is a simple substitution cipher where each letter is
// replaced by the number of the order they're in. For example, A corresponds to
// 1, B = 2, C = 3, etc..
//
// author(s): [Focusucof](https://github.com/Focusucof)
// see a1z26_test.go for tests

// ref: https://www.dcode.fr/letter-number-cipher
package a1z26

import (
	"strconv"
	"strings"
)

var encryptionMap = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// helper function to find the index of a given value in an array
func indexOfLetter(slice []string, item string) string {
	for i := range slice {
		if slice[i] == item {
			return strconv.Itoa(i + 1)
		}
	}
	return ""
}

func Encrypt(input string) string {
	var result string

	input = strings.ToLower(input)
	for _, char := range input {
		if string(char) != " " {
			result += indexOfLetter(encryptionMap, string(char)) // replace letter with the corresponding number
			result += "-"                                        // separate letters in the same word by dashes
		} else {
			result = result[:len(result)-1] // remove remove dash and replace with a space
			result += " "
		}

	}
	result = result[:len(result)-1] // remove leading dash at end of string
	return result
}

func Decrypt(input string) string {
	var result string = ""

	words := strings.Fields(input) // split string into multiple words

	for _, word := range words {
		word = strings.ReplaceAll(word, "-", " ") // replace dashes with spaces

		letters := strings.Fields(word) // split letters into a slice
		for _, i := range letters {
			index, _ := strconv.Atoi(string(i)) // convert char to int so we can use it as an array index
			result += encryptionMap[index-1]
		}
		result += " "
	}
	result = result[:len(result)-1] // remove leading space at end of string
	return result
}
