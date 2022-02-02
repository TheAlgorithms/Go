package longestincreasingsubsequencegreedy

import (
	"testing"
)

func TestLongestIncreasingSubsequenceGreedy(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "Empty slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "Longest increasing subsequence is 1",
			slice:    []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Longest increasing subsequence is 6",
			slice:    []int{1, 2, 3, 4, 5, 6, 6},
			expected: 6,
		},
		{
			name:     "Longest increasing subsequence is 4",
			slice:    []int{2, 4, 3, 7, 4, 5},
			expected: 4,
		},
	}

	for _, test := range testCases {
		if LongestIncreasingSubsequenceGreedy(test.slice) != test.expected {
			t.Fatalf("Failed for: %s", test.name)
		}
	}
}
