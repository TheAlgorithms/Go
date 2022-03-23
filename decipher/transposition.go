package decipher

import (
	"sort"
	"strings"
)

// getKey is a duplicate of the method present in the cipher package - this is
// duplicated to make sure these packages are independent.
func getKey(keyWord string) []int {
	keyWord = strings.ToLower(keyWord)
	word := []rune(keyWord)
	var sortedWord = make([]rune, len(word))
	copy(sortedWord, word)
	sort.Slice(sortedWord, func(i, j int) bool { return sortedWord[i] < sortedWord[j] })
	usedLettersMap := make(map[rune]int)
	wordLength := len(word)
	resultKey := make([]int, wordLength)
	for i := 0; i < wordLength; i++ {
		char := word[i]
		numberOfUsage := usedLettersMap[char]
		resultKey[i] = getIndex(sortedWord, char) + numberOfUsage + 1 //+1 -so that indexing does not start at 0
		numberOfUsage++
		usedLettersMap[char] = numberOfUsage
	}
	return resultKey
}

// getIndex is a function that is a duplicate of the function that is present in the cipher package.
// We have duplicated it, in order to make sure that these two packages remain independent.
func getIndex(wordSet []rune, subString rune) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}
	return 0
}

// Transposition deciphers the text encrypted using the Transposition cipher.
func Transposition(text []rune, keyWord string) (string, error) {
	key := getKey(keyWord)
	textLength := len(text)
	if textLength <= 0 {
		return "", NoTextToDecryptError
	}
	keyLength := len(key)
	if keyLength <= 0 {
		return "", KeyMissingError
	}
	space := ' '
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, space)
	}
	result := ""
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result += string(transposition)
	}
	return result, nil
}
