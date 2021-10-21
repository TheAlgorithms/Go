package coloring

func (g *Graph) TryBipartiteColoring() map[int]Color {
	// 0 is uncolored, 1/2 is colors
	colors := make(map[int]Color)
	visited := make(map[int]bool)

	for i := range g.edges {
		colors[i] = 0
		visited[i] = false
	}

	var dfs func(int)
	dfs = func(s int) {
		if !visited[s] {
			visited[s] = true
			coloring := []Color{0, 2, 1}

			for n := range g.edges[s] {
				if colors[n] == 0 {
					colors[n] = coloring[colors[s]]
				}
				dfs(n)
			}
		}
	}

	for i := range g.edges {
		if colors[i] == 0 {
			colors[i] = 1
			dfs(i)
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
