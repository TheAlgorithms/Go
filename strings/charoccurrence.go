// charoccurrence.go
// description: Chars occurrence in the sentence
// details: An algorithm to count chars in string
// author(s) [Moein](https://github.com/mo1ein)
// see charoccurrence_test.go

package strings

func Count(text string) map[rune]int {
	charOccurrence := make(map[rune]int, 0)
	for _, c := range text {
		if _, ok := charOccurrence[c]; !ok {
			charOccurrence[c] = 0
		}
		charOccurrence[c]++
	}
	return charOccurrence
}
