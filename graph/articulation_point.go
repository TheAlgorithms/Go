package graph

import "github.com/TheAlgorithms/Go/math/min"

func (graph *Graph) ArticulationPoint() []bool {
	var (
		is_ap = make([]bool, graph.vertices) // boolean array to mark articulation points
		visited = make([]bool, graph.vertices) // boolean array to track whether a vertex is visited or not
		child = make([]int, graph.vertices) // array to store the number of children of a vertex
		discovery = make([]int, graph.vertices) // array to store the time of discovery of a vertex by DFS
		earliest_discovery = make([]int, graph.vertices) // array to store the earliest discovered vertex reachable from a vertex
		time = 0 // time variable to keep track of the time of discovery
	)
	
	func dfs(graph *Graph, u int, parent int) {
		visited[u] = true

		discovery[u] = time // Mark the time of discovery of a vertex
		earliest_discovery[u] = discovery[u] // set the earliest discovery time to the discovered time
		time++ // increment the time

		for v, _ := range graph.edges[u] {
			if v == parent {
				continue
			}
			if !visited[v] {
				child[u]++
				dfs(graph, v, u)
				earliest_discovery[u] = min.Min(earliest_discovery[u], earliest_discovery[v])
				if earliest_discovery[v] >= discovery[u] {
					is_ap[u] = true
				}
			} else {
				earliest_discovery[u] = min.Min(earliest_discovery[u], discovery[v])
			}
		}
	}

	dfs(graph, 0, -1)

	if child[0] == 1 {
		is_ap[0] = false
	}

	return is_ap
}