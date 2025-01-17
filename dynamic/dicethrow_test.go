package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseDiceThrow struct {
	numDice   int
	numFaces  int
	targetSum int
	expected  int
}

// getDiceThrowTestCases provides the test cases for DiceThrow
func getDiceThrowTestCases() []testCaseDiceThrow {
	return []testCaseDiceThrow{
		{2, 6, 7, 6},  // Two dice, six faces each, sum = 7
		{1, 6, 3, 1},  // One die, six faces, sum = 3
		{3, 4, 5, 6},  // Three dice, four faces each, sum = 5
		{1, 6, 1, 1},  // One die, six faces, sum = 1
		{2, 6, 12, 1}, // Two dice, six faces each, sum = 12
		{3, 6, 18, 1}, // Three dice, six faces each, sum = 18
		{2, 6, 20, 0}, // Two dice, six faces each, sum = 20 (impossible)
		{1, 1, 1, 1},  // One die, one face, sum = 1
		{1, 1, 2, 0},  // One die, one face, sum = 2 (impossible)
		{2, 1, 2, 1},  // Two dice, one face each, sum = 2
	}
}

// TestDiceThrow tests the DiceThrow function with basic test cases
func TestDiceThrow(t *testing.T) {
	t.Run("Basic test cases", func(t *testing.T) {
		for _, tc := range getDiceThrowTestCases() {
			actual := dynamic.DiceThrow(tc.numDice, tc.numFaces, tc.targetSum)
			if actual != tc.expected {
				t.Errorf("DiceThrow(%d, %d, %d) = %d; expected %d", tc.numDice, tc.numFaces, tc.targetSum, actual, tc.expected)
			}
		}
	})
}
