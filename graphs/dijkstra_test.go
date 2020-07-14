package graph

import (
	"fmt"
	"testing"
)

func TestGraph_Dijkstra(t *testing.T) {

	graph := Graph{}
	graph.InsertEdge(1, 2, 4)
	graph.InsertEdge(2, 1, 4)
	graph.InsertEdge(1, 8, 8)
	graph.InsertEdge(8, 1, 8)
	graph.InsertEdge(2, 3, 8)
	graph.InsertEdge(3, 2, 8)
	graph.InsertEdge(2, 8, 11)
	graph.InsertEdge(8, 2, 11)
	graph.InsertEdge(3, 4, 7)
	graph.InsertEdge(4, 3, 7)
	graph.InsertEdge(3, 6, 4)
	graph.InsertEdge(6, 3, 4)
	graph.InsertEdge(3, 9, 2)
	graph.InsertEdge(9, 3, 2)
	graph.InsertEdge(4, 5, 9)
	graph.InsertEdge(5, 4, 9)
	graph.InsertEdge(4, 6, 14)
	graph.InsertEdge(6, 4, 14)
	graph.InsertEdge(5, 6, 10)
	graph.InsertEdge(6, 5, 10)
	graph.InsertEdge(6, 7, 2)
	graph.InsertEdge(7, 6, 2)
	graph.InsertEdge(7, 8, 1)
	graph.InsertEdge(8, 7, 1)
	graph.InsertEdge(8, 9, 7)
	graph.InsertEdge(9, 8, 7)
	shortestPaths := Dijkstra(&graph, 1)
	fmt.Println(graph.String())
	for node, cost := range shortestPaths {
		fmt.Printf("Distance from %d to %d = %d\n", 1, node, cost)
	}

}
