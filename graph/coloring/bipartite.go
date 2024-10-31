package coloring

// Bipartite.go
// description: Implementation of the Bipartite graph coloring algorithm
// details: A bipartite graph is a graph whose vertices can be divided into two disjoint sets U and V such that every edge connects a vertex in U to one in V. The Bipartite graph coloring algorithm is used to determine if a graph is bipartite or not.
// time complexity: O(V+E) where V is the number of vertices and E is the number of edges in the graph
// space complexity: O(V) where V is the number of vertices in the graph

func (g *Graph) TryBipartiteColoring() map[int]Color {
	// 0 is uncolored, 1/2 are colors
	colors := make(map[int]Color)
	visited := make(map[int]bool)

	for i := range g.edges {
		colors[i] = 0
		visited[i] = false
	}

	var colorNode func(int)
	colorNode = func(s int) {
		visited[s] = true
		coloring := []Color{0, 2, 1}

		for n := range g.edges[s] {
			if colors[n] == 0 {
				colors[n] = coloring[colors[s]]
			}
			if !visited[n] {
				colorNode(n)
			}
		}
	}

	for i := range g.edges {
		if colors[i] == 0 {
			colors[i] = 1
			colorNode(i)
		}
	}

	return colors
}

// basically tries to color the graph in two colors if each edge
// connects 2 differently colored nodes the graph can be considered bipartite
func BipartiteCheck(N int, edges [][]int) bool {
	var graph Graph
	for i := 0; i < N; i++ {
		graph.AddVertex(i)
	}
	for _, e := range edges {
		graph.AddEdge(e[0], e[1])
	}
	return graph.ValidateColorsOfVertex(graph.TryBipartiteColoring()) == nil
}
