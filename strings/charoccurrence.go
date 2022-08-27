// charoccurrence.go
// description: Chars occurrence in the sentence
// details: An algorithm to count chars in string
// author(s) [Moein](https://github.com/mo1ein)
// see charoccurrence_test.go

package strings

// CountChars count chars occurrence from string
// It returns the map of rune:int
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
