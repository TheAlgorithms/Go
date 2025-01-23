package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseMaxSubArraySum struct {
	nums     []int
	expected int
}

func getMaxSubArraySumTestCases() []testCaseMaxSubArraySum {
	return []testCaseMaxSubArraySum{
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},   // Kadane's algorithm example
		{[]int{-1, -2, -3, -4}, -1},               // All negative numbers, max single element
		{[]int{5, 4, -1, 7, 8}, 23},               // Positive numbers with a large sum
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // Mixed with a maximum subarray of length 4
		{[]int{1, 2, 3, 4, 5}, 15},                // All positive numbers, sum is the entire array
		{[]int{-1, -2, -3, -4, -5}, -1},           // Only negative numbers, largest single element
		{[]int{0, 0, 0, 0, 0}, 0},                 // Array of zeros, maximum subarray is zero
		{[]int{3}, 3},                             // Single positive number
		{[]int{-1}, -1},                           // Single negative number
	}
}

func TestMaxSubArraySum(t *testing.T) {
	t.Run("Max SubArray Sum test cases", func(t *testing.T) {
		for _, tc := range getMaxSubArraySumTestCases() {
			actual := dynamic.MaxSubArraySum(tc.nums)
			if actual != tc.expected {
				t.Errorf("MaxSubArraySum(%v) = %v; expected %v", tc.nums, actual, tc.expected)
			}
		}
	})
}
