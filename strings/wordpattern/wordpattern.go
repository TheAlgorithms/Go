// wordpattern.go
// description: Word Pattern
// details:
// Given a pattern and a string s, find if s follows the same pattern.
// example:
// Input: pattern = "xyyx", s = "go kotlin kotlin go"
// Output: true
// author(s) [red_byte](https://github.com/i-redbyte)
// see wordpattern_test.go

package wordpattern

import (
	"strings"
)

func WordPattern(s, pattern string) bool {
	words := strings.Split(s, " ")
	n := len(words)
	if n != len(pattern) {
		return false
	}
	wordsMap := make(map[string]string)
	patternsMap := make(map[string]string)
	for i, v := range strings.Split(pattern, "") {
		if w, ok := wordsMap[v]; ok && w != words[i] {
			return false
		} else if _, ok := patternsMap[words[i]]; ok && patternsMap[words[i]] != v {
			return false
		}
		wordsMap[v] = words[i]
		patternsMap[words[i]] = v
	}
	return true
}
