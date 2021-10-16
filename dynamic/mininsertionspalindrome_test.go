package dynamic

import (
	"fmt"
	"testing"
)

func Test_FindMinInsertionsPalindrome(t *testing.T) {

	var testCases = []struct {
		str      string
		expected int
	}{
		{"a", 0},
		{"ab", 1},
		{"abcd", 3},
		{"abcda", 2},
		{"abcde", 4},
		{"geeks", 3},
		{"thealgorithms", 10},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("string: %s", testCases[i].str), func(t *testing.T) {
			result := FindMinInsertionsPalindrome(testCases[i].str)
			if result != testCases[i].expected {
				t.Errorf("Expected %d insertions for %s, but found %d", testCases[i].expected, testCases[i].str, result)
			}
		})
	}
}
