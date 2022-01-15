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
	charArray := make([]int, 256) //This array is used to represent the ASCII values.
	for _, char := range cleanText1 {
		charArray[char]++ //Increments the index for the characters in cleanText1 Example: charArray['a']++ => charArray[97]++
	}
	for _, char := range cleanText2 {
		charArray[char]-- //Decrements the index for the characters in cleanText2  Example: charArray['a']-- => charArray[97]--
	}
	for i := 0; i < 256; i++ {
		if charArray[i] != 0 { //If the characters in both texts are balanced, it is an anagram otherwise No
			return false
		}
	}
	return true

}
