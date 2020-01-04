/*
KMP is an algorithm for finding the position of pattern in a text, and the time complexity is O(n + m) where n is the length of text, and m is the length of pattern.
*/

package main

import "fmt"

var next []int

func GetNext(s string) {
	next = make([]int, len(s)+1)
	next[0] = 0
	next[1] = 0

	for i := 1; i < len(s); i++ {
		j := next[i]
		for j > 0 && s[j] != s[i] {
			j = next[j]
		}

		if s[j] == s[i] {
			next[i+1] = j + 1
		} else {
			next[i+1] = 0
		}
	}
}

func Match(text, pattern string) int {
	var j int = 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = next[j]
		}

		if pattern[j] == text[i] {
			j++

			if j == len(pattern) {
				return i - j
			}
		}
	}

	return -1
}

func main() {
	text, pattern := "AAAAAAAAAAB", "AAAAB"

	GetNext(pattern)
	fmt.Println("The next array of pattern is:")
	fmt.Println(next)

	fmt.Println("The first matching position is:")
	fmt.Println(Match(text, pattern))
}
