/*
This algorithm tries to search the given pattern in the given text.
If pattern is found from index position i in the text,
it is added to positions.

Time Complexity : O(n*m)
    n = length of text
    m = length of pattern
*/

package search

func Naive(text string, pattern string) []int {
	var positions []int
	for i := 0; i <= len(text)-len(pattern); i++ {
		var match bool = true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}

		}
		if match {
			positions = append(positions, i)
		}
	}
	return positions
}
