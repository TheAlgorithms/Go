package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

var longestCommonSubsequenceExamples = []struct {
	stringA  string
	stringB  string
	expected int
}{
	{"ABCDGH", "AEDFHR", 3},
	{"AGGTAB", "GXTXAYB", 4},
	{"programming", "gaming", 6},
	{"physics", "smartphone", 2},
	{"computer", "food", 1},
	{"123", "12345", 3},
	{"XYZ", "XYZ", 3},
	{"XYZ", "XYZa", 3},
	{"XYZ", "aXYZ", 3},
	{"0123", "abc", 0},
	{"abcdef", "aXbXcXXXdeXXf", 6},
	{"", "abc", 0},
	{"", "", 0},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tc := range longestCommonSubsequenceExamples {
		actual := dynamic.LongestCommonSubsequence(
			tc.stringA, tc.stringB,
			len(tc.stringA), len(tc.stringB))
		if actual != tc.expected {
			t.Errorf(
				"LongestCommonSubsequence of %s and %s expected to be: %d, but got: %d",
				tc.stringA, tc.stringB, tc.expected, actual)
		}
	}
}

func TestLongestCommonSubsequenceIsSymmetric(t *testing.T) {
	for _, tc := range longestCommonSubsequenceExamples {
		actual := dynamic.LongestCommonSubsequence(
			tc.stringB, tc.stringA,
			len(tc.stringB), len(tc.stringA))
		if actual != tc.expected {
			t.Errorf(
				"LongestCommonSubsequence of %s and %s expected to be: %d, but got: %d",
				tc.stringB, tc.stringA, tc.expected, actual)
		}
	}
}
