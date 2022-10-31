package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseLCS struct {
	stringA  string
	stringB  string
	expected int
}

func getLCSTestCases() []testCaseLCS {
	return []testCaseLCS{
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
		{"££", "££", 2},
		{"x笑x笑", "aaa笑a笑", 2},
		{"xYxY", "aaaYaY", 2},
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		for _, tc := range getLCSTestCases() {
			actual := dynamic.LongestCommonSubsequence(tc.stringA, tc.stringB)
			if actual != tc.expected {
				t.Errorf("expected: %d, but got: %d", tc.expected, actual)
			}
		}
	})

	t.Run("Symmetry test", func(t *testing.T) {
		for _, tc := range getLCSTestCases() {
			actual := dynamic.LongestCommonSubsequence(tc.stringB, tc.stringA)
			if actual != tc.expected {
				t.Errorf("expected: %d, but got: %d", tc.expected, actual)
			}
		}
	})
}
