// topological.go
// description: Topological sort
// details: Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u v, vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.
// time complexity: O(V+E) where V is the number of vertices and E is the number of edges in the graph
// space complexity: O(V) where V is the number of vertices in the graph
// reference: https://en.wikipedia.org/wiki/Topological_sorting

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
