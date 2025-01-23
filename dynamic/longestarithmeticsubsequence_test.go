package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseLongestArithmeticSubsequence struct {
	nums     []int
	expected int
}

func getLongestArithmeticSubsequenceTestCases() []testCaseLongestArithmeticSubsequence {
	return []testCaseLongestArithmeticSubsequence{
		{[]int{3, 6, 9, 12}, 4},                  // Arithmetic sequence of length 4
		{[]int{9, 4, 7, 2, 10}, 3},               // Arithmetic sequence of length 3
		{[]int{20, 1, 15, 3, 10, 5, 8}, 4},       // Arithmetic sequence of length 4
		{[]int{1, 2, 3, 4, 5}, 5},                // Arithmetic sequence of length 5
		{[]int{10, 7, 4, 1}, 4},                  // Arithmetic sequence of length 4
		{[]int{1, 5, 7, 8, 5, 3, 4, 3, 1, 2}, 4}, // Arithmetic sequence of length 4
		{[]int{1, 3, 5, 7, 9}, 5},                // Arithmetic sequence of length 5
		{[]int{5, 10, 15, 20}, 4},                // Arithmetic sequence of length 4
		{[]int{1}, 1},                            // Single element, length is 1
		{[]int{}, 0},                             // Empty array, length is 0
	}
}

func TestLongestArithmeticSubsequence(t *testing.T) {
	t.Run("Longest Arithmetic Subsequence test cases", func(t *testing.T) {
		for _, tc := range getLongestArithmeticSubsequenceTestCases() {
			actual := dynamic.LongestArithmeticSubsequence(tc.nums)
			if actual != tc.expected {
				t.Errorf("LongestArithmeticSubsequence(%v) = %v; expected %v", tc.nums, actual, tc.expected)
			}
		}
	})
}
