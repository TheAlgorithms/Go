package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseBurstBalloons struct {
	nums     []int
	expected int
}

func getBurstBalloonsTestCases() []testCaseBurstBalloons {
	return []testCaseBurstBalloons{
		{[]int{3, 1, 5, 8}, 167}, // Maximum coins from [3,1,5,8]
		{[]int{1, 5}, 10},        // Maximum coins from [1,5]
		{[]int{1}, 1},            // Single balloon
		{[]int{}, 0},             // No balloons
	}

}

func TestMaxCoins(t *testing.T) {
	t.Run("Burst Balloons test cases", func(t *testing.T) {
		for _, tc := range getBurstBalloonsTestCases() {
			actual := dynamic.MaxCoins(tc.nums)
			if actual != tc.expected {
				t.Errorf("MaxCoins(%v) = %d; expected %d", tc.nums, actual, tc.expected)
			}
		}
	})
}
