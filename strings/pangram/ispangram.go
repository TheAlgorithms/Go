// ispangram.go
// description: Checks if a given string is pangram or not
// details: A pangram is a sentence or expression that uses all the letters of the alphabet.
// Reference: https://www.geeksforgeeks.org/pangram-checking/
// Author : Kavitha J

package pangram

import (
	"regexp"
	"strings"
)

func cleanString(text string) string {
	cleanText := strings.ToLower(text)                      // Convert to lowercase
	cleanText = strings.Join(strings.Fields(cleanText), "") // Remove spaces
	regex, _ := regexp.Compile(`[^\p{L}\p{N} ]+`)           // Regular expression for alphanumeric only characters
	return regex.ReplaceAllString(cleanText, "")
}

func IsPangram(text string) bool {
	cleanText := cleanString(text)
	if len(cleanText) < 26 {
		return false
	}
	var data = make(map[rune]bool)
	for _, i := range cleanText {
		data[i] = true
	}
	return len(data) == 26
}
