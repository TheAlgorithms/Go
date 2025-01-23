package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseEggDropping struct {
	eggs     int
	floors   int
	expected int
}

func getEggDroppingTestCases() []testCaseEggDropping {
	return []testCaseEggDropping{
		{1, 10, 10}, // One egg, need to test all floors
		{2, 10, 4},  // Two eggs and ten floors
		{3, 14, 4},  // Three eggs and fourteen floors
		{2, 36, 8},  // Two eggs and thirty-six floors
		{2, 0, 0},   // Two eggs, zero floors
	}

}

func TestEggDropping(t *testing.T) {
	t.Run("Egg Dropping test cases", func(t *testing.T) {
		for _, tc := range getEggDroppingTestCases() {
			actual := dynamic.EggDropping(tc.eggs, tc.floors)
			if actual != tc.expected {
				t.Errorf("EggDropping(%d, %d) = %d; expected %d", tc.eggs, tc.floors, actual, tc.expected)
			}
		}
	})
}
