package main

import "testing"

func TestNaivePatternSearch(t *testing.T) {
	text := "ABAAABCDBBABCDDEBCABC"
	pattern := "ABC"

	var positions []int = naivePatternSearch(text, pattern)

	if len(positions) > 3 {
		t.Errorf("test NaivePatternSearch failed")
	}
}
