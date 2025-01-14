package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseOptimalBST struct {
	keys     []int
	freq     []int
	n        int
	expected int
}

func getOptimalBSTTestCases() []testCaseOptimalBST {
	return []testCaseOptimalBST{
		{[]int{10, 12, 20}, []int{34, 8, 50}, 3, 142},                  // Example with 3 keys
		{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}, 5, 300}, // Example with 5 keys
		{[]int{10, 20}, []int{5, 10}, 2, 15},                           // Simple case with 2 keys
		{[]int{10}, []int{100}, 1, 100},                                // Single key case
		{[]int{1, 2, 3, 4}, []int{10, 100, 20, 50}, 4, 180},            // Case with 4 keys
		{[]int{1, 3, 5, 7, 9}, []int{10, 50, 30, 40, 20}, 5, 230},      // Case with 5 keys
	}
}

func TestOptimalBST(t *testing.T) {
	t.Run("Optimal Binary Search Tree test cases", func(t *testing.T) {
		for _, tc := range getOptimalBSTTestCases() {
			actual := dynamic.OptimalBST(tc.keys, tc.freq, tc.n)
			if actual != tc.expected {
				t.Errorf("OptimalBST(%v, %v, %d) = %d; expected %d", tc.keys, tc.freq, tc.n, actual, tc.expected)
			}
		}
	})
}
