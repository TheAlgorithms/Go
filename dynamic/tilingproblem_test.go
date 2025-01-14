package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseTilingProblem struct {
	n        int
	expected int
}

func getTilingProblemTestCases() []testCaseTilingProblem {
	return []testCaseTilingProblem{
		{1, 1},   // Base case: 1 way to tile a 2x1 grid
		{2, 2},   // 2 ways to tile a 2x2 grid
		{3, 3},   // 3 ways to tile a 2x3 grid
		{4, 5},   // 5 ways to tile a 2x4 grid
		{5, 8},   // 8 ways to tile a 2x5 grid
		{6, 13},  // 13 ways to tile a 2x6 grid
		{10, 89}, // 89 ways to tile a 2x10 grid
		{0, 1},   // Edge case: 1 way to tile a 2x0 grid (no tiles)
		{7, 21},  // 21 ways to tile a 2x7 grid
		{8, 34},  // 34 ways to tile a 2x8 grid
	}
}

func TestTilingProblem(t *testing.T) {
	t.Run("Tiling Problem test cases", func(t *testing.T) {
		for _, tc := range getTilingProblemTestCases() {
			actual := dynamic.TilingProblem(tc.n)
			if actual != tc.expected {
				t.Errorf("TilingProblem(%d) = %d; expected %d", tc.n, actual, tc.expected)
			}
		}
	})
}
