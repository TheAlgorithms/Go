package graph

// Assumes that graph given is valid and possible to
// get a topo ordering.
// constraints are array of []int{a, b}, representing
// an edge going from a to b
func TopoSort(N int, constraints [][]int) []int {
	dependencies := make([]int, N)
	edges := make(map[int][]int)
	for _, c := range constraints {
		a := c[0]
		b := c[1]
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

			for _, v := range edges[s0] {
				if !visited[v] {
					visited[v] = true
					stack = append(stack, v)
				}
			}
		}
	}

	return ans
}
