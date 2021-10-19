package graph

func addto(adj map[int][]int, a int, b int) {
	_, isin := adj[a]
	if !isin {
		adj[a] = []int{}
	}
	adj[a] = append(adj[a], b)
}

// basically tries to color the graph in two colors if each edge
// connects 2 differently colored nodes the graph can be considered bipartite
func IsBipartite(N int, edges [][]int) bool {
	adj := map[int][]int{}
	for _, e := range edges {
		addto(adj, e[0], e[1])
		addto(adj, e[1], e[0])
	}

	// 0 is uncolored, 1/2 is colors
	colors := make([]int, N)
	colors[0] = 1

	stack := []int{0}
	isgood := true

walk:
	// O(n) algorithm to walk the tree
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		alt := 1
		if colors[s] == 1 {
			alt = 2
		}

		// color the neighbours and all them to check if they are not
		// in the stack yet
		for _, n := range adj[s] {
			// because uncolored=0, unvisited neighbours will always
			// be safe as our current node will always be colored
			if colors[n] == colors[s] {
				isgood = false
				break walk
			}

			if colors[n] == 0 {
				colors[n] = alt
				stack = append(stack, n)
			}
		}
	}

	return isgood
}
