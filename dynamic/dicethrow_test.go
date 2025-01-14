package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseDiceThrow struct {
	m, n, sum int
	expected  int
}

func getDiceThrowTestCases() []testCaseDiceThrow {
	return []testCaseDiceThrow{
		{2, 6, 7, 6},   // Two dice, six faces each, sum = 7
		{1, 6, 3, 1},   // One die, six faces, sum = 3
		{3, 4, 5, 6},   // Three dice, four faces each, sum = 5
		{2, 4, 8, 5},   // Two dice, four faces each, sum = 8
		{1, 6, 1, 1},   // One die, six faces, sum = 1
		{2, 6, 12, 1},  // Two dice, six faces each, sum = 12
		{3, 6, 18, 1},  // Three dice, six faces each, sum = 18
		{2, 6, 20, 0},  // Two dice, six faces each, sum = 20 (impossible)
		{3, 8, 15, 51}, // Three dice, eight faces each, sum = 15
		{1, 1, 1, 1},   // One die, one face, sum = 1
		{1, 1, 2, 0},   // One die, one face, sum = 2 (impossible)
		{2, 1, 2, 1},   // Two dice, one face each, sum = 2
	}
}

func TestDiceThrow(t *testing.T) {
	t.Run("Basic test cases", func(t *testing.T) {
		for _, tc := range getDiceThrowTestCases() {
			actual := dynamic.DiceThrow(tc.m, tc.n, tc.sum)
			if actual != tc.expected {
				t.Errorf("DiceThrow(%d, %d, %d) = %d; expected %d", tc.m, tc.n, tc.sum, actual, tc.expected)
			}
		}
	})
}
