package graph

import "testing"

func TestHasCycle(t *testing.T) {
	graph := Graph{Directed: true}
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {4, 0}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	if !graph.HasCycle() {
		t.Error("answer of hasCycle is not correct")
	}

	graph = Graph{Directed: true}
	edges = [][]int{{0, 1}, {1, 2}, {2, 6}, {4, 0}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	if graph.HasCycle() {
		t.Error("answer of hasCycle is not correct")
	}
}
