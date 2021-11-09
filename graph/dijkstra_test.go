package graph

import (
	"testing"
)

var tc_dijkstra = []struct {
	name     string
	edges    [][]int
	node0    int
	node1    int
	expected int
}{
	{
		"straight line graph",
		[][]int{{0, 1, 5}, {1, 2, 2}},
		0, 2, 7,
	},
	{
		"unconnected node",
		[][]int{{0, 1, 5}},
		0, 2, -1,
	},
	{
		"double paths",
		[][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}},
		0, 3, 9,
	},
	{
		"double paths extended",
		[][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}, {3, 4, 1}},
		0, 4, 10,
	},
}

func TestDijkstra(t *testing.T) {
	for _, tc := range tc_dijkstra {
		t.Run(tc.name, func(t *testing.T) {
			var graph Graph
			for _, edge := range tc.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			actual, _ := graph.Dijkstra(tc.node0, tc.node1)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d, from node %d to %d, with %v",
					tc.expected, actual, tc.node0, tc.node1, tc.edges)
			}
		})
	}
}
