// This file contains the graph coloring implementation using backtracking
// time complexity: O(V^V) where V is the number of vertices in the graph
// space complexity: O(V) where V is the number of vertices in the graph
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring

// ColorUsingBacktracking will return the Color of each vertex and the
// total number of different colors used, using backtracking
func (g *Graph) ColorUsingBacktracking() (map[int]Color, int) {
	vertexColors := make(map[int]Color, g.vertices)
	g.colorVertex(0, vertexColors)

	colorsUsed := 0
	for _, cr := range vertexColors {
		if colorsUsed < int(cr) {
			colorsUsed = int(cr)
		}
	}
	return vertexColors, colorsUsed
}

// colorVertex will try to color provided vertex, v
func (g *Graph) colorVertex(v int, color map[int]Color) bool {
	// If all vertices are colored, the colors store will be completely filled.
	if len(color) == g.vertices {
		return true
	}

	// As the upper bound of no. of colors is the no. of vertices in graph,
	// try assigning each color to the vertex v
	for cr := Color(1); cr <= Color(g.vertices); cr++ {
		// Use the color, cr for vertex, v if it is safe to use, by
		// checking its neighbours
		safe := true
		for nb := range g.edges[v] {
			// cr, color is not safe if color of nb, crnb is not equal to cr
			if crnb, ok := color[nb]; ok && crnb == cr {
				safe = false
				break
			}
		}
		if safe {
			color[v] = cr
			if g.colorVertex(v+1, color) {
				return true
			}
			delete(color, v)
		}
	}
	return false
}
