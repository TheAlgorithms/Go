// Kahn's algorithm computes a topological ordering of a directed acyclic graph (DAG).
// Time Complexity: O(V + E)
// Space Complexity: O(V + E)
// Reference: https://en.wikipedia.org/wiki/Topological_sorting#Kahn's_algorithm
// see graph.go, topological.go, kahn_test.go

package graph

// Kahn's algorithm computes a topological ordering of a directed acyclic graph (DAG).
// `n` is the number of vertices,
// `dependencies` is a list of directed edges, where each pair [a, b] represents
// a directed edge from a to b (i.e. b depends on a).
// Vertices are assumed to be labelled 0, 1, ..., n-1.
// If the graph is not a DAG, the function returns nil.
func Kahn(n int, dependencies [][]int) []int {
	g := Graph{vertices: n, Directed: true}
	// track the in-degree (number of incoming edges) of each vertex
	inDegree := make([]int, n)

	// populate g with edges, increase the in-degree counts accordingly
	for _, d := range dependencies {
		// make sure we don't add the same edge twice
		if _, ok := g.edges[d[0]][d[1]]; !ok {
			g.AddEdge(d[0], d[1])
			inDegree[d[1]]++
		}
	}

	// queue holds all vertices with in-degree 0
	// these vertices have no dependency and thus can be ordered first
	queue := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// order holds a valid topological order
	order := make([]int, 0, n)

	// process the dependency-free vertices
	// every time we process a vertex, we "remove" it from the graph
	for len(queue) > 0 {
		// pop the first vertex from the queue
		vtx := queue[0]
		queue = queue[1:]
		// add the vertex to the topological order
		order = append(order, vtx)
		// "remove" all the edges coming out of this vertex
		// every time we remove an edge, the corresponding in-degree reduces by 1
		// if all dependencies on a vertex is removed, enqueue the vertex
		for neighbour := range g.edges[vtx] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	// if the graph is a DAG, order should contain all the certices
	if len(order) != n {
		return nil
	}
	return order
}
