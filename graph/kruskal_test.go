package graph

import (
	"fmt"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	// Define test cases with inputs, expected outputs, and sample graphs
	var testCases = []struct {
		n     int
		graph []Edge
		cost  int
	}{
		// Test Case 1
		{
			n: 5,
			graph: []Edge{
				{Start: 0, End: 1, Weight: 4},
				{Start: 0, End: 2, Weight: 13},
				{Start: 0, End: 3, Weight: 7},
				{Start: 0, End: 4, Weight: 7},
				{Start: 1, End: 2, Weight: 9},
				{Start: 1, End: 3, Weight: 3},
				{Start: 1, End: 4, Weight: 7},
				{Start: 2, End: 3, Weight: 10},
				{Start: 2, End: 4, Weight: 14},
				{Start: 3, End: 4, Weight: 4},
			},
			cost: 20,
		},
		// Test Case 2
		{
			n: 3,
			graph: []Edge{
				{Start: 0, End: 1, Weight: 12},
				{Start: 0, End: 2, Weight: 18},
				{Start: 1, End: 2, Weight: 6},
			},
			cost: 18,
		},
		// Test Case 3
		{
			n: 4,
			graph: []Edge{
				{Start: 0, End: 1, Weight: 2},
				{Start: 0, End: 2, Weight: 1},
				{Start: 0, End: 3, Weight: 2},
				{Start: 1, End: 2, Weight: 1},
				{Start: 1, End: 3, Weight: 2},
				{Start: 2, End: 3, Weight: 3},
			},
			cost: 4,
		},
		// Test Case 4
		{
			n: 2,
			graph: []Edge{
				{Start: 0, End: 1, Weight: 4000000},
			},
			cost: 4000000,
		},
		// Test Case 5
		{
			n: 1,
			graph: []Edge{
				{Start: 0, End: 0, Weight: 0},
			},
			cost: 0,
		},
	}

	// Iterate through the test cases and run the tests
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			// Execute KruskalMST for the current test case
			_, computed := KruskalMST(testCase.n, testCase.graph)

			// Compare the computed result with the expected result
			if computed != testCase.cost {
				t.Errorf("Test Case %d, Expected: %d, Computed: %d", i, testCase.cost, computed)
			}
		})
	}
}
