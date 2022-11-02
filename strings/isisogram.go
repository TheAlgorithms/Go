// Checks if a given string is an isogram.
// A first-order isogram is a word in which no letter of the alphabet occurs more than once.
// A second-order isogram is a word in which each letter appears twice.
// A third-order isogram is a word in which each letter appears three times.
// wiki: https://en.wikipedia.org/wiki/Heterogram_(literature)#Isograms
// Author: M3talM0nk3y

package strings

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
			letters[l] += 1

			if letters[l] > 3 {
				return false
			}

			continue
		}
		letters[l] = 1
	}

	mapVals := make(map[int]bool)
	for _, v := range letters {
		mapVals[v] = true
	}

	return len(mapVals) == 1
}
