package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestClimbingStairs(t *testing.T) {
	td := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{10, 89},
		{15, 987},
		{20, 10946},
		{-1, 0},
		{-5, 0},
	}

	for _, tc := range td {
		name := fmt.Sprintf("ClimbingStairs(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			actual := dynamic.ClimbingStairs(tc.n)

			if actual != tc.expected {
				t.Errorf(
					"ClimbingStairs(%d) = %d, expected %d",
					tc.n,
					actual,
					tc.expected,
				)
			}
		})
	}
}

func ExampleClimbingStairs() {
	fmt.Print(dynamic.ClimbingStairs(4))
	// Output: 5
}
