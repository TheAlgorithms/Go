/*
This algorithm tries to search the given pattern in the given text.
If pattern is found from index position i in the text,
it is added to positions.

Time Complexity : O(n*m)
    n = length of text
    m = length of pattern
*/

package main

import (
	"fmt"
)

func naivePatternSearch(text string, pattern string) []int {
	var positions []int
	// added the "=" sign to the condition in for loop i to include the last value which return 18
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

func main() {
	text := "ABAAABCDBBABCDDEBCABC"
	pattern := "ABC"
	var positions []int = naivePatternSearch(text, pattern)
	if len(positions) == 0 {
		fmt.Printf("Pattern not found in given text!")
	} else {
		fmt.Printf("Pattern found in following position:\n")
		fmt.Printf("%v", positions)
	}
}
