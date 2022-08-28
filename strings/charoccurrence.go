// charoccurrence.go
// description: An algorithm which counts the number
// of times a character occurred in a string.
// author(s) [Moein](https://github.com/mo1ein)
// see charoccurrence_test.go

package strings

// CountChars counts the number of a times a character
// has occurred in the provided string argument and
// returns a map with `rune` as keys and the count as value.
func CountChars(text string) map[rune]int {
	charMap := make(map[rune]int, 0)
	for _, c := range text {
		if _, ok := charMap[c]; !ok {
			charMap[c] = 0
		}
		charMap[c]++
	}
	return charMap
}
