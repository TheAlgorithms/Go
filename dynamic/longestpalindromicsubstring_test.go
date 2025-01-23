package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseLongestPalindromicSubstring struct {
	s        string
	expected string
}

func getLongestPalindromicSubstringTestCases() []testCaseLongestPalindromicSubstring {
	return []testCaseLongestPalindromicSubstring{
		{"babad", "bab"},                   // Example with multiple palindromes
		{"cbbd", "bb"},                     // Example with longest even palindrome
		{"a", "a"},                         // Single character, palindrome is itself
		{"", ""},                           // Empty string, no palindrome
		{"racecar", "racecar"},             // Whole string is a palindrome
		{"abcba", "abcba"},                 // Palindrome in the middle
		{"aabbcc", "aa"},                   // Multiple substrings, longest "aa"
		{"madam", "madam"},                 // Full palindrome string
		{"forgeeksskeegfor", "geeksskeeg"}, // Complex palindrome in the middle
	}
}

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Run("Longest Palindromic Substring test cases", func(t *testing.T) {
		for _, tc := range getLongestPalindromicSubstringTestCases() {
			actual := dynamic.LongestPalindromicSubstring(tc.s)
			if actual != tc.expected {
				t.Errorf("LongestPalindromicSubstring(%q) = %q; expected %q", tc.s, actual, tc.expected)
			}
		}
	})
}
