/*isanagram.go
description: Checks if two given strings are anagram or not
details: An anagram is a word or phrase that's formed by rearranging the letters of another word or phrase
Reference: https://www.geeksforgeeks.org/pangram-checking/
Author : Venkatasubramanian */

package anagram

import "strings"

func cleanString(text string) []rune {
	cleanText := strings.Join(strings.Fields(text), "") // Remove spaces
	return []rune(cleanText)
}

func isanagram(text1 string, text2 string) bool {
	cleanText1 := cleanString(text1)
	cleanText2 := cleanString(text2)
	if len(cleanText1) != len(cleanText2) {
		return false
	}
	charMap := make(map[rune]int)

	for i := 0; i < len(cleanText1); i++ {
		if charMap[cleanText1[i]] != 0 {
			charMap[cleanText1[i]]++
		} else {
			charMap[cleanText1[i]] = 1
		}
	}
	for i := 0; i < len(cleanText2); i++ {
		if charMap[cleanText2[i]] > 0 {
			charMap[cleanText2[i]]--
		} else {
			return false
		}
	}

	return true

}
