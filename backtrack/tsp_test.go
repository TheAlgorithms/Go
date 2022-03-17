package backtrack

import (
	"reflect"
	"testing"
)

var testCasesForTsp = []struct {
	name         string
	input        [][]int
	expectedCost int
	expectedPath []int
}{
	{
		name: "four cities",
		input: [][]int{
			{0, 30, 6, 4},
			{30, 0, 5, 10},
			{6, 5, 0, 20},
			{4, 10, 20, 0},
		},
		expectedCost: 25,
		expectedPath: []int{0, 3, 1, 2},
	},
	{
		name: "three cities",
		input: [][]int{
			{0, 10, 5},
			{10, 0, 10},
			{5, 10, 0},
		},
		expectedCost: 25, // 5 + 10 + 10
		expectedPath: []int{0, 2, 1},
	},
	{
		name: "two cities",
		input: [][]int{
			{0, 10},
			{10, 0},
		},
		expectedCost: 20, // 10 + 10
		expectedPath: []int{0, 1},
	},
}

func TestTSP(t *testing.T) {
	for _, testCase := range testCasesForTsp {
		t.Run(testCase.name, func(t *testing.T) {
			actualCost, actualPath := TSP(testCase.input)
			if actualCost != testCase.expectedCost {
				t.Errorf("expected cost %d, but got %d", testCase.expectedCost, actualCost)
			}
			if len(actualPath) != len(testCase.expectedPath) {
				t.Errorf("expected path length %v, but got %v", len(testCase.expectedPath), len(actualPath))
			}
			if reflect.DeepEqual(actualPath, testCase.expectedPath) == false {
				t.Errorf("expected path %v, but got %v", testCase.expectedPath, actualPath)
			}
		})
	}
}
