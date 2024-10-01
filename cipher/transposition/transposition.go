// transposition.go
// description: Transposition cipher
// details:
// Implementation "Transposition cipher" is a method of encryption by which the positions held by units of plaintext (which are commonly characters or groups of characters) are shifted according to a regular system, so that the ciphertext constitutes a permutation of the plaintext [Transposition cipher](https://en.wikipedia.org/wiki/Transposition_cipher)
// time complexity: O(n)
// space complexity: O(n)
// author(s) [red_byte](https://github.com/i-redbyte)
// see transposition_test.go

package transposition

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var ErrNoTextToEncrypt = errors.New("no text to encrypt")
var ErrKeyMissing = errors.New("missing Key")

const placeholder = ' '

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

func getIndex(wordSet []rune, subString rune) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}
	return 0
}

func Encrypt(text []rune, keyWord string) ([]rune, error) {
	key := getKey(keyWord)
	keyLength := len(key)
	textLength := len(text)
	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}
	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}
	if text[len(text)-1] == placeholder {
		return nil, fmt.Errorf("%w: cannot encrypt a text, %q, ending with the placeholder char %q", ErrNoTextToEncrypt, text, placeholder)
	}
	n := textLength % keyLength

	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}
	textLength = len(text)
	var result []rune
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[key[j]-1] = text[i+j]
		}
		result = append(result, transposition...)
	}
	return result, nil
}

func Decrypt(text []rune, keyWord string) ([]rune, error) {
	key := getKey(keyWord)
	textLength := len(text)
	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}
	keyLength := len(key)
	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}
	var result []rune
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result = append(result, transposition...)
	}
	result = []rune(strings.TrimRight(string(result), string(placeholder)))
	return result, nil
}
