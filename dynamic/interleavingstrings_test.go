package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseInterleaving struct {
	s1, s2, s3 string
	expected   bool
}

func getInterleavingTestCases() []testCaseInterleaving {
	return []testCaseInterleaving{
		{"aab", "axy", "aaxaby", true},   // Valid interleaving
		{"aab", "axy", "abaaxy", false},  // Invalid interleaving
		{"", "", "", true},               // All empty strings
		{"abc", "", "abc", true},         // Only s1 matches s3
		{"", "xyz", "xyz", true},         // Only s2 matches s3
		{"abc", "xyz", "abxcyz", true},   // Valid interleaving
		{"aaa", "aaa", "aaaaaa", true},   // Identical strings
		{"aaa", "aaa", "aaaaaaa", false}, // Extra character
		{"abc", "def", "abcdef", true},   // Concatenation order
		{"abc", "def", "adbcef", true},   // Valid mixed interleaving
	}
}

func TestIsInterleave(t *testing.T) {
	t.Run("Interleaving Strings test cases", func(t *testing.T) {
		for _, tc := range getInterleavingTestCases() {
			actual := dynamic.IsInterleave(tc.s1, tc.s2, tc.s3)
			if actual != tc.expected {
				t.Errorf("IsInterleave(%q, %q, %q) = %v; expected %v", tc.s1, tc.s2, tc.s3, actual, tc.expected)
			}
		}
	})
}
