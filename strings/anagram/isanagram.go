// isanagram.go
// description: Checks if two given strings are anagram or not
// details: An anagram is a word or phrase that's formed by rearranging the letters of another word or phrase
// Reference: https://www.geeksforgeeks.org/pangram-checking/
// Author : Venkatasubramanian

package anagram

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

func isanagram(text1 string, text2 string) bool {
	cleanText1 := cleanString(text1)
	cleanText2 := cleanString(text2)
	if len(cleanText1) != len(cleanText2) {
		return false
	}
	arr := make([]int, 256)
	for _, char := range cleanText1 {
		arr[char]++
	}
	for _, char := range cleanText2 {
		arr[char]--
	}
	for i := 0; i < 256; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true

}
