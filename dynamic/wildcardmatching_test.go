package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseWildcardMatching struct {
	s        string
	p        string
	expected bool
}

func getWildcardMatchingTestCases() []testCaseWildcardMatching {
	return []testCaseWildcardMatching{
		{"aa", "a*", true},                    // * can match zero or more characters
		{"mississippi", "m??*ss*?i*pi", true}, // Complex pattern with multiple wildcards
		{"aa", "a", false},                    // No match due to no wildcard
		{"ab", "?*", true},                    // ? matches any single character, * matches remaining
		{"abcd", "a*d", true},                 // * matches the characters between 'a' and 'd'
		{"abcd", "a*c", false},                // No match as 'c' doesn't match the last character 'd'
		{"abc", "*", true},                    // * matches the entire string
		{"abc", "a*c", true},                  // * matches 'b'
		{"abc", "a?c", true},                  // ? matches 'b'
		{"abc", "a?d", false},                 // ? cannot match 'd'
		{"", "a*", true},                      // * can match an empty string
		{"", "", true},                        // Both strings empty, so they match
		{"a", "?", true},                      // ? matches any single character
		{"a", "*", true},                      // * matches any number of characters, including one
	}
}

func TestIsMatch(t *testing.T) {
	t.Run("Wildcard Matching test cases", func(t *testing.T) {
		for _, tc := range getWildcardMatchingTestCases() {
			actual := dynamic.IsMatch(tc.s, tc.p)
			if actual != tc.expected {
				t.Errorf("IsMatch(%q, %q) = %v; expected %v", tc.s, tc.p, actual, tc.expected)
			}
		}
	})
}
