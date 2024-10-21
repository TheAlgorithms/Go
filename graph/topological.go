package graph

// Topological assumes that graph given is valid and that its
// possible to get a topological ordering.
// constraints are array of []int{a, b}, representing
// an edge going from a to b
func Topological(N int, constraints [][]int) []int {
	dependencies := make([]int, N)
	nodes := make([]int, N)
	for i := range nodes {
		nodes[i] = i
	}
	edges := make([][]bool, N)
	for i := range edges {
		edges[i] = make([]bool, N)
	}

	for _, c := range constraints {
		a := c[0]
		b := c[1]
		dependencies[b]++
		edges[a][b] = true
	}

	var answer []int
	for s := 0; s < N; s++ {
		// Only start walking from top level nodes
		if dependencies[s] == 0 {
			route, _ := DepthFirstSearchHelper(s, N, nodes, edges, true)
			answer = append(answer, route...)
		}
	}

	return answer
}
