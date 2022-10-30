// Checks if a given string is an isogram.
// An isogram is a word, phrase, or sentence in which no letter of the alphabet occurs more than once.
// wiki: https://en.wikipedia.org/wiki/Heterogram_(literature)#Isograms
// Author: M3talM0nk3y

package isogram

import (
	"regexp"
	"strings"
)

func hasDigit(text string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(text)
}

func hasSymbol(text string) bool {
	re := regexp.MustCompile(`[-!@#$%^&*()+]`)
	return re.MatchString(text)
}

func IsIsogram(text string) bool {
	text = strings.ToLower(text)
	text = strings.Join(strings.Fields(text), "")

	if hasDigit(text) || hasSymbol(text) {
		return false
	}

	letters := make(map[string]int)
	for _, c := range text {
		l := string(c)
		if _, ok := letters[l]; ok {
			return false
		}
		letters[l] = 1
	}

	return len(letters) == len(text)
}
