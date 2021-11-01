// longestpalindrome.go
// description: Manacher's algorithm (Longest palindromic substring)
// details:
// An algorithm with linear running time that allows you to get compressed information about all palindromic substrings of a given string. - [Manacher's algorithm](https://en.wikipedia.org/wiki/Longest_palindromic_substring)
// author(s) [red_byte](https://github.com/i-redbyte)
// see longestpalindrome_test.go

package manacher

import (
	"github.com/TheAlgorithms/Go/math/min"
	"strings"
)

func makeBoundaries(s string) string {
	var result strings.Builder
	result.WriteRune('#')
	for _, ch := range s {
		if ch != ' ' { //ignore space as palindrome character
			result.WriteRune(ch)
		}
		result.WriteRune('#')
	}
	return result.String()
}

func nextBoundary(s string) string {
	var result strings.Builder
	for _, ch := range s {
		if ch != '#' {
			result.WriteRune(ch)
		}
	}
	return result.String()
}

func LongestPalindrome(s string) string {
	boundaries := makeBoundaries(s)
	b := make([]int, len(boundaries))
	k := 0
	index := 0
	maxLen := 0
	maxCenterSize := 0
	for i := range b {
		if i < k {
			b[i] = min.Int(b[2*index-i], k-i)
		} else {
			b[i] = 1
		}
		for i-b[i] >= 0 && i+b[i] < len(boundaries) && boundaries[i-b[i]] == boundaries[i+b[i]] {
			b[i] += 1
		}
		if maxLen < b[i]-1 {
			maxLen = b[i] - 1
			maxCenterSize = i
		}
		if b[i]+i-1 > k {
			k = b[i] + i - 1
			index = i
		}
	}
	return nextBoundary(boundaries[maxCenterSize-maxLen : maxCenterSize+maxLen])
}
