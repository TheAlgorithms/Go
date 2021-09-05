package graph

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var bfsTestData = []struct {
		description string
		start       int
		end         int
		nodes       int
		edges       [][]int
		expected1   bool
		expected2   int
	}{
		{
			"test 1 connected with distance 2",
			0,
			5,
			6,
			[][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 1},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 1, 0, 0, 0, 0},
			},
			true,
			2,
		},
		{
			"test 2 connected with distance 4",
			0,
			5,
			6,
			[][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 1, 0},
			},
			true,
			4,
		},
		{
			"test 2 not connected",
			0,
			5,
			6,
			[][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			false,
			0,
		},
	}
	for _, test := range bfsTestData {
		t.Run(test.description, func(t *testing.T) {
			r1, r2 := BreadthFirstSearch(test.start, test.end,
				test.nodes, test.edges)
			if r1 != test.expected1 || r2 != test.expected2 {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Nodes '%v' and Edges '%v' start from '%d' and end in '%d' "+
					"was expecting '%v' with distance '%d' but result was '%v','%d'",
					test.nodes, test.edges, test.start, test.end, test.expected1, test.expected2, r1, r2)
			}
		})
	}
}
