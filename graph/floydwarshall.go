// Floyd-Warshall algorithm
// time complexity: O(V^3) where V is the number of vertices in the graph
// space complexity: O(V^2) where V is the number of vertices in the graph
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

package graph

import "math"

// WeightedGraph defining matrix to use 2d array easier
type WeightedGraph [][]float64

// Defining maximum value. If two vertices share this value, it means they are not connected
var Inf = math.Inf(1)

// FloydWarshall Returns all pair's shortest path using Floyd Warshall algorithm
func FloydWarshall(graph WeightedGraph) WeightedGraph {
	// If graph is empty, returns nil
	if len(graph) == 0 || len(graph) != len(graph[0]) {
		return nil
	}

	for i := 0; i < len(graph); i++ {
		//If graph matrix width is different than the height, returns nil
		if len(graph[i]) != len(graph) {
			return nil
		}
	}

	numVertices := len(graph)

	// Initializing result matrix and filling it up with same values as given graph
	result := make(WeightedGraph, numVertices)

	for i := 0; i < numVertices; i++ {
		result[i] = make([]float64, numVertices)
		for j := 0; j < numVertices; j++ {
			result[i][j] = graph[i][j]
		}
	}

	// Running over the result matrix and following the algorithm
	for k := 0; k < numVertices; k++ {
		for i := 0; i < numVertices; i++ {
			for j := 0; j < numVertices; j++ {
				// If there is a less costly path from i to j node, remembering it
				if result[i][j] > result[i][k]+result[k][j] {
					result[i][j] = result[i][k] + result[k][j]
				}
			}
		}
	}

	return result
}
