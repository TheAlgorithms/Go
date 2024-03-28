// Edmond-Karp algorithm is an implementation of the Ford-Fulkerson method
// to compute max-flow between a pair of source-sink vertices in a weighted graph
// It uses BFS (Breadth First Search) to find the residual paths
// Time Complexity: O(V * E^2) where V is the number of vertices and E is the number of edges
// Space Complexity: O(V + E) Because we keep residual graph in size of the original graph
// Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, and Clifford Stein. 2009. Introduction to Algorithms, Third Edition (3rd. ed.). The MIT Press.

package graph

import (
	"math"
)

// Returns a mapping of vertices as path, if there is any from source to sink
// Otherwise, returns nil
func FindPath(rGraph WeightedGraph, source int, sink int) map[int]int {
	queue := make([]int, 0)
	marked := make([]bool, len(rGraph))
	marked[source] = true
	queue = append(queue, source)
	parent := make(map[int]int)

	// BFS loop with saving the path found
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for i := 0; i < len(rGraph[v]); i++ {
			if !marked[i] && rGraph[v][i] > 0 {
				parent[i] = v
				// Terminate the BFS, if we reach to sink
				if i == sink {
					return parent
				}
				marked[i] = true
				queue = append(queue, i)
			}
		}
	}
	// source and sink are not in the same connected component
	return nil
}

func EdmondKarp(graph WeightedGraph, source int, sink int) float64 {
	// Check graph emptiness
	if len(graph) == 0 {
		return 0.0
	}

	// Check correct dimensions of the graph slice
	for i := 0; i < len(graph); i++ {
		if len(graph[i]) != len(graph) {
			return 0.0
		}
	}

	rGraph := make(WeightedGraph, len(graph))
	for i := 0; i < len(graph); i++ {
		rGraph[i] = make([]float64, len(graph))
	}
	// Init the residual graph with the same capacities as the original graph
	copy(rGraph, graph)

	maxFlow := 0.0

	for {
		parent := FindPath(rGraph, source, sink)
		if parent == nil {
			break
		}
		// Finding the max flow over the path returned by BFS
		// i.e. finding minimum residual capacity amonth the path edges
		pathFlow := math.MaxFloat64
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			if rGraph[u][v] < pathFlow {
				pathFlow = rGraph[u][v]
			}
		}

		// update residual capacities of the edges and
		// reverse edges along the path
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			rGraph[u][v] -= pathFlow
			rGraph[v][u] += pathFlow
		}

		// Update the total flow found so far
		maxFlow += pathFlow
	}

	return maxFlow
}
