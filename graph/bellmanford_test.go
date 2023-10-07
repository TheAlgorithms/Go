package graph

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestBellmanford(t *testing.T) {

	var testCases = []struct {
		name        string
		edges       [][]int
		vertices    int
		start       int
		end         int
		isReachable bool
		distance    int
		err         error
	}{
		{
			"single edge",
			[][]int{
				{0, 1, 1},
			},
			2, 0, 1, true, 1, nil,
		},
		{
			"negative weights",
			[][]int{
				{0, 1, 1},
				{1, 2, -3},
				{2, 1, 4},
				{2, 3, 1},
			},
			4, 0, 1, true, 1, nil,
		},
		{
			"negative cycle",
			[][]int{
				{0, 1, 1},
				{1, 2, -3},
				{2, 1, 1},
				{2, 3, 1},
			},
			4, 0, 1, false, -1, errors.New("negative weight cycle present"),
		},
		{
			"unreachable vertex",
			[][]int{
				{0, 6, 771},
				{0, 9, 782},
				{1, 2, 454},
				{2, 8, 48},
				{3, 8, 249},
				{3, 9, 880},
				{3, 5, 280},
				{7, 1, 92},
				{7, 2, 497},
				{8, 1, 102},
				{8, 4, 977},
			},
			10, 8, 3, false, int(math.Inf(1)), nil,
		},
		{
			"disconnected graph",
			[][]int{
				{0, 1, 10},
				{2, 3, 15},
				{3, 5, 10},
			},
			6, 0, 3, false, int(math.Inf(1)), nil,
		},
		{
			"multiple paths",
			[][]int{
				{0, 1, 5},
				{1, 2, 10},
				{1, 3, 30},
				{2, 4, 10},
				{4, 5, 15},
				{3, 5, 10},
			},
			6, 0, 5, true, 40, nil,
		},
		{
			"random 1",
			[][]int{
				{0, 1, 10},
				{1, 2, 10},
				{0, 2, 100},
				{2, 0, -10},
				{1, 2, 1},
			},
			3, 0, 1, true, 10, nil,
		},
		{
			"random 2",
			[][]int{
				{0, 1, 5498},
				{2, 0, 7679},
				{0, 3, 4999},
				{1, 2, 8629},
				{1, 3, -948},
				{2, 3, 6231},
			},
			4, 0, 3, true, 4550, nil,
		},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			// Initializing graph, adding edges
			graph := New(test.vertices)
			graph.Directed = true
			for _, edge := range test.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			resIsReachable, resDistance, resError := graph.BellmanFord(test.start, test.end)
			if resDistance != test.distance {
				t.Errorf("Distance, Expected: %d, Computed: %d", test.distance, resDistance)
			}
			if resIsReachable != test.isReachable {
				t.Errorf("Reachable, Expected: %t, Computed: %t", test.isReachable, resIsReachable)
			}
			if !errors.Is(test.err, resError) {
				if resError == nil || test.err == nil {
					t.Errorf("Reachable, Expected: %s, Computed: %s", test.err, resError)
				} else if resError.Error() != test.err.Error() {
					t.Errorf("Reachable, Expected: %s, Computed: %s", test.err.Error(), resError.Error())
				}
			}
		})
	}
}
