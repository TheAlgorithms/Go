// This file contains the graph coloring implementation using Greedy Approach.
// time complexity: O(V^2) where V is the number of vertices in the graph
// space complexity: O(V) where V is the number of vertices in the graph
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring

import "sort"

// ColorUsingGreedyApproach will return the Color of each vertex and the
// total number of different colors used, using a greedy approach, based on
// the number of edges (or degree) from any vertex.
func (g *Graph) ColorUsingGreedyApproach() (map[int]Color, int) {
	degreeOfVertex := make([]struct{ degree, vertex int }, 0, g.vertices)
	for v, neighbours := range g.edges {
		degreeOfVertex = append(degreeOfVertex,
			struct{ degree, vertex int }{
				vertex: v,
				degree: len(neighbours),
			},
		)
	}
	// sort the degreeOfVertex in decreasing order of degrees
	sort.Slice(degreeOfVertex, func(i, j int) bool {
		return degreeOfVertex[i].degree > degreeOfVertex[j].degree
	})

	vertexColors := make(map[int]Color, g.vertices)
	// Start with a color and assign the color to all possible vertices in the degreeOfVertex slice
	// and then, re-iterate with new color for all those which are left
	for color := 1; color <= g.vertices; color++ {
	vertexLoop:
		for _, val := range degreeOfVertex {
			// skip, if already assigned
			if _, ok := vertexColors[val.vertex]; ok {
				continue vertexLoop
			}
			// Check its neighbours
			for ng := range g.edges[val.vertex] {
				if vertexColors[ng] == Color(color) {
					// not possible to use this color for val.vertex
					continue vertexLoop
				}
			}
			// Assign color to the vertex
			vertexColors[val.vertex] = Color(color)
		}
		// continue till all the vertices are colored
		if len(vertexColors) == g.vertices {
			return vertexColors, color
		}
	}
	return vertexColors, g.vertices
}
