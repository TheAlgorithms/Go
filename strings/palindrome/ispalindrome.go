package palindrome

import (
	"regexp"
	"strings"
)

func CleanString(text string) string {
	clean_text := strings.ToLower(text)
	clean_text = strings.Join(strings.Fields(clean_text), "") // Remove spaces
	regex, _ := regexp.Compile("[^a-zA-Z0-9 ]+")              // Regular expression for alphanumeric only characters
	return regex.ReplaceAllString(clean_text, "")
}

func IsPalindrome(text string) bool {
	clean_text := CleanString(text)
	var i, j int
	for i = 0; i < len(clean_text)/2; i++ {
		j = len(clean_text) - 1 - i
		if clean_text[i] != clean_text[j] {
			return false
		}
	}
	return true
}
