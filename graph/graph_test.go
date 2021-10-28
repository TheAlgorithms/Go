// Tests for directed and undirected graphs

package graph

import (
	"fmt"
	"testing"
)

var graphTestCases = []struct {
	name     string
	edges    [][]int
	vertices int
}{
	{
		"single edge",
		[][]int{
			{0, 1, 1},
		},
		2,
	},
	{
		"many edges",
		[][]int{
			{0, 1, 1},
			{0, 2, 2},
			{1, 3, 4},
			{3, 4, 3},
			{4, 8, 3},
			{4, 9, 1},
			{7, 8, 2},
			{8, 9, 2},
		},
		10,
	},
	{
		"cycles",
		[][]int{
			{0, 1, 1},
			{0, 2, 2},
			{1, 3, 4},
			{3, 4, 3},
			{4, 2, 1},
		},
		5,
	},
	{
		"disconnected graphs",
		[][]int{
			{0, 1, 5},
			{2, 4, 5},
			{3, 8, 5},
		},
		2,
	},
}

func TestDirectedGraph(t *testing.T) {

	// Testing self-loops separately only for directed graphs.
	// For undirected graphs each edge already creates a self-loop.
	directedGraphTestCases := append(graphTestCases, struct {
		name     string
		edges    [][]int
		vertices int
	}{
		"self-loops",
		[][]int{
			{0, 1, 1},
			{1, 2, 2},
			{2, 1, 3},
		},
		3,
	})

	for _, test := range directedGraphTestCases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			// Initializing graph, adding edges
			graph := New(test.vertices)
			graph.Directed = true
			for _, edge := range test.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			if graph.vertices != test.vertices {
				t.Errorf("Number of vertices, Expected: %d, Computed: %d", test.vertices, graph.vertices)
			}
			edgeCount := 0
			for _, e := range graph.edges {
				edgeCount += len(e)
			}
			if edgeCount != len(test.edges) {
				t.Errorf("Number of edges, Expected: %d, Computed: %d", len(test.edges), edgeCount)
			}
			for _, edge := range test.edges {
				if val, found := graph.edges[edge[0]][edge[1]]; !found || val != edge[2] {
					t.Errorf("Edge {%d->%d (%d)} not found", edge[0], edge[1], edge[2])
				}
			}
		})
	}
}

func TestUndirectedGraph(t *testing.T) {

	for _, test := range graphTestCases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			// Initializing graph, adding edges
			graph := New(test.vertices)
			for _, edge := range test.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			if graph.vertices != test.vertices {
				t.Errorf("Number of vertices, Expected: %d, Computed: %d", test.vertices, graph.vertices)
			}
			edgeCount := 0
			for _, e := range graph.edges {
				edgeCount += len(e)
			}
			if edgeCount != len(test.edges)*2 {
				t.Errorf("Number of edges, Expected: %d, Computed: %d", len(test.edges)*2, edgeCount)
			}
			for _, edge := range test.edges {
				if val, found := graph.edges[edge[0]][edge[1]]; !found || val != edge[2] {
					t.Errorf("Edge {%d->%d (%d)} not found", edge[0], edge[1], edge[2])
				}
			}
		})
	}
}
