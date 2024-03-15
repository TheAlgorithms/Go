package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimMST(t *testing.T) {

	var testCases = []struct {
		edges    []Edge
		vertices int
		start    int
		cost     int
		mst      []Edge
	}{
		{
			edges: []Edge{
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
			vertices: 5,
			start:    0,
			cost:     20,
			mst: []Edge{
				{Start: 0, End: 1, Weight: 4},
				{Start: 1, End: 3, Weight: 3},
				{Start: 3, End: 4, Weight: 4},
				{Start: 1, End: 2, Weight: 9},
			},
		},
		{
			edges: []Edge{
				{Start: 0, End: 1, Weight: 4},
				{Start: 0, End: 7, Weight: 8},
				{Start: 1, End: 2, Weight: 8},
				{Start: 1, End: 7, Weight: 11},
				{Start: 2, End: 3, Weight: 7},
				{Start: 2, End: 5, Weight: 4},
				{Start: 2, End: 8, Weight: 2},
				{Start: 3, End: 4, Weight: 9},
				{Start: 3, End: 5, Weight: 14},
				{Start: 4, End: 5, Weight: 10},
				{Start: 5, End: 6, Weight: 2},
				{Start: 6, End: 7, Weight: 1},
				{Start: 6, End: 8, Weight: 6},
				{Start: 7, End: 8, Weight: 7},
			},
			vertices: 9,
			start:    3,
			cost:     37,
			mst: []Edge{
				{Start: 3, End: 2, Weight: 7},
				{Start: 2, End: 8, Weight: 2},
				{Start: 2, End: 5, Weight: 4},
				{Start: 5, End: 6, Weight: 2},
				{Start: 6, End: 7, Weight: 1},
				{Start: 2, End: 1, Weight: 8},
				{Start: 1, End: 0, Weight: 4},
				{Start: 3, End: 4, Weight: 9},
			},
		},
		{
			edges: []Edge{
				{Start: 0, End: 1, Weight: 2},
				{Start: 0, End: 3, Weight: 6},
				{Start: 1, End: 2, Weight: 3},
				{Start: 1, End: 3, Weight: 8},
				{Start: 1, End: 4, Weight: 5},
				{Start: 2, End: 4, Weight: 7},
				{Start: 3, End: 4, Weight: 9},
			},
			vertices: 5,
			start:    2,
			cost:     16,
			mst: []Edge{
				{Start: 2, End: 1, Weight: 3},
				{Start: 1, End: 0, Weight: 2},
				{Start: 1, End: 4, Weight: 5},
				{Start: 0, End: 3, Weight: 6},
			},
		},
		{
			edges: []Edge{
				{Start: 0, End: 0, Weight: 0},
			},
			vertices: 1,
			start:    0,
			cost:     0,
			mst:      nil,
		},
		{
			edges: []Edge{
				{Start: 0, End: 1, Weight: 1},
				{Start: 0, End: 2, Weight: 6},
				{Start: 0, End: 3, Weight: 5},
				{Start: 1, End: 2, Weight: 2},
				{Start: 1, End: 4, Weight: 4},
				{Start: 2, End: 4, Weight: 9},
			},
			vertices: 5,
			start:    4,
			cost:     12,
			mst: []Edge{
				{Start: 4, End: 1, Weight: 4},
				{Start: 1, End: 0, Weight: 1},
				{Start: 1, End: 2, Weight: 2},
				{Start: 0, End: 3, Weight: 5},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			// Initializing graph, adding edges
			graph := New(testCase.vertices)
			graph.Directed = false
			for _, edge := range testCase.edges {
				graph.AddWeightedEdge(int(edge.Start), int(edge.End), edge.Weight)
			}

			computedMST, computedCost := graph.PrimMST(Vertex(testCase.start))

			// Compare the computed result with the expected result
			if computedCost != testCase.cost {
				t.Errorf("Test Case %d, Expected Cost: %d, Computed: %d", i, testCase.cost, computedCost)
			}

			if !reflect.DeepEqual(testCase.mst, computedMST) {
				t.Errorf("Test Case %d, Expected MST: %v, Computed: %v", i, testCase.mst, computedMST)
			}
		})
	}
}
