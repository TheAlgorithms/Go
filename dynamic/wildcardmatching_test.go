package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

// testCaseWildcardMatching holds the test cases for the Wildcard Matching problem
type testCaseWildcardMatching struct {
	s        string
	p        string
	expected bool
}

// getWildcardMatchingTestCases returns a list of test cases for the Wildcard Matching problem
func getWildcardMatchingTestCases() []testCaseWildcardMatching {
	return []testCaseWildcardMatching{
		{"aa", "a*", true},     // '*' can match zero or more characters
		{"aa", "a", false},     // No match due to no wildcard
		{"ab", "?*", true},     // '?' matches any single character, '*' matches remaining
		{"abcd", "a*d", true},  // '*' matches the characters between 'a' and 'd'
		{"abcd", "a*c", false}, // No match as 'c' doesn't match the last character 'd'
		{"abc", "*", true},     // '*' matches the entire string
		{"abc", "a*c", true},   // '*' matches 'b'
		{"abc", "a?c", true},   // '?' matches 'b'
		{"abc", "a?d", false},  // '?' cannot match 'd'
		{"", "", true},         // Both strings empty, so they match
		{"a", "?", true},       // '?' matches any single character
		{"a", "*", true},       // '*' matches any number of characters, including one
	}
}

// TestIsMatch tests the IsMatch function with various test cases
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
