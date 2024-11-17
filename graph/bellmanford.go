// The Bellman–Ford algorithm is an algorithm that computes shortest paths from a
// single source vertex to all of the other vertices in a weighted directed graph.
// It is slower than Dijkstra but capable of handling negative edge weights.
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
// Implementation is based on the book 'Introduction to Algorithms' (CLRS)
// time complexity: O(V*E) where V is the number of vertices and E is the number of edges in the graph
// space complexity: O(V) where V is the number of vertices in the graph

package graph

import (
	"errors"
	"math"
)

func (g *Graph) BellmanFord(start, end int) (isReachable bool, distance int, err error) {
	INF := math.Inf(1)
	distances := make([]float64, g.vertices)

	// Set all vertices to unreachable, initialize source
	for i := 0; i < g.vertices; i++ {
		distances[i] = INF
	}
	distances[start] = 0

	// Making iterations equal to #vertices
	for n := 0; n < g.vertices; n++ {

		// Looping over all edges
		for u, adjacents := range g.edges {
			for v, weightUV := range adjacents {

				// If new shorter distance is found, update distance value (relaxation step)
				if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
					distances[v] = newDistance
				}
			}
		}
	}

	// Check for negative weight cycle
	for u, adjacents := range g.edges {
		for v, weightUV := range adjacents {
			if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
				return false, -1, errors.New("negative weight cycle present")
			}
		}
	}

	return distances[end] != INF, int(distances[end]), nil
}
