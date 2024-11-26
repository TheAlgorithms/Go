// Package graph provides algorithms to analyze graph structures.
package graph

import "github.com/TheAlgorithms/Go/math/min"

// apHelper stores auxiliary data used to identify articulation points in a graph.
type apHelper struct {
	isAP              []bool
	visited           []bool
	childCount        []int
	discoveryTime     []int
	earliestDiscovery []int
}

// ArticulationPoint identifies articulation points in a graph. It returns a boolean slice
// where each element indicates whether a vertex is an articulation point.
// Worst Case Time Complexity: O(|V| + |E|)
// Auxiliary Space: O(|V|)
// Reference: https://en.wikipedia.org/wiki/Biconnected_component and https://cptalks.quora.com/Cut-Vertex-Articulation-point
func ArticulationPoint(graph *Graph) []bool {
	// Time variable to keep track of the discovery time of a vertex
	time := 0

	// Initialize apHelper instance with the required data structures
	apHelperInstance := &apHelper{
		isAP:              make([]bool, graph.vertices),
		visited:           make([]bool, graph.vertices),
		childCount:        make([]int, graph.vertices),
		discoveryTime:     make([]int, graph.vertices),
		earliestDiscovery: make([]int, graph.vertices),
	}

	// Start traversal from the root (0)
	articulationPointHelper(apHelperInstance, 0, -1, &time, graph)

	// Check if the root has only one child, making it non-articulate
	if apHelperInstance.childCount[0] == 1 {
		apHelperInstance.isAP[0] = false
	}

	return apHelperInstance.isAP
}

// articulationPointHelper recursively traverses the graph using DFS and marks articulation points.
// It updates `childCount`, `discoveryTime`, and `earliestDiscovery` slices for the given vertex.
func articulationPointHelper(
	apHelperInstance *apHelper,
	vertex int,
	parent int,
	time *int,
	graph *Graph,
) {
	apHelperInstance.visited[vertex] = true

	// Set discovery and earliest discovery times for the vertex
	apHelperInstance.discoveryTime[vertex] = *time
	apHelperInstance.earliestDiscovery[vertex] = *time
	*time++

	for nextVertex := range graph.edges[vertex] {
		if nextVertex == parent {
			continue
		}

		if apHelperInstance.visited[nextVertex] {
			// Update the earliest discovery time to the smallest reachable discovery time
			apHelperInstance.earliestDiscovery[vertex] = min.Int(
				apHelperInstance.earliestDiscovery[vertex],
				apHelperInstance.discoveryTime[nextVertex],
			)
			continue
		}

		// Increment child count and perform recursive traversal for DFS
		apHelperInstance.childCount[vertex]++
		articulationPointHelper(apHelperInstance, nextVertex, vertex, time, graph)

		// Update the earliest discovery time post DFS
		apHelperInstance.earliestDiscovery[vertex] = min.Int(
			apHelperInstance.earliestDiscovery[vertex],
			apHelperInstance.earliestDiscovery[nextVertex],
		)

		// Mark vertex as articulation point if condition meets
		if apHelperInstance.earliestDiscovery[nextVertex] >= apHelperInstance.discoveryTime[vertex] {
			apHelperInstance.isAP[vertex] = true
		}
	}
}
