package mergealternately

import (
	"fmt"
)

// Return the alternative arrangements of the two strings.

func MergeAlternately(word1 string, word2 string) string {
	ans := ""

	t := [...]string{word1, word2}

	if len(t[0]) > len(t[1]) {
		t[0], t[1] = t[1], t[0]
	}

	for i := 0; i < len(t[0]); i++ {
		ans += fmt.Sprintf("%c%c", word1[i], word2[i])
	}

	ans += t[1][len(t[0]):]

	return ans
}
