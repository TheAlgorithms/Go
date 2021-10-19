package graph

// Assumes that graph given is valid
// and possible to get a topo ordering
func TopoSort(N int, constraints [][]int) []int {
	dependencies := make([]int, N)
	edges := make(map[int][]int)
	for i := 0; i < len(constraints); i++ {
		a := constraints[i][0]
		b := constraints[i][1]
		dependencies[b]++

		_, isin := edges[a]
		if !isin {
			edges[a] = []int{}
		}
		edges[a] = append(edges[a], b)
	}

	ans := []int{}
	stack := []int{}
	visited := make(map[int]bool, N)

	// O(n) algorithm for walking the
	// tree in topo order
	for s := 0; s < N; s++ {
		// Only start walking from top level nodes
		if dependencies[s] > 0 {
			continue
		}

		// dfs walk down in topo order
		visited[s] = true
		stack := append(stack, s)
		for len(stack) > 0 {
			s0 := stack[len(stack)-1]
			ans = append(ans, s0)
			stack = stack[:len(stack)-1]

			for v := 0; v < len(edges[s0]); v++ {
				v0 := edges[s0][v]
				if !visited[v0] {
					visited[v0] = true
					stack = append(stack, v0)
				}
			}
		}
	}

	return ans
}
