// This file contains the simple structural implementation of undirected
// graph. Used in coloring algorithms [here](./coloring)
// Author(s): [Shivam](https://github.com/Shivam010)

package graph

// UndirectedGraph provides a structure to store an undirected graph.
// It is safe to use its empty object.
type UndirectedGraph struct {
	vertices int
	edges    map[int]map[int]struct{}
}

// AddVertex will add a new vertex in the graph, if the vertex already
// exist it will do nothing
func (g *UndirectedGraph) AddVertex(v int) {
	if g.edges == nil {
		g.edges = make(map[int]map[int]struct{})
	}

	// Check if vertex is present or not
	if _, ok := g.edges[v]; !ok {
		g.vertices++
		g.edges[v] = make(map[int]struct{})
	}
}

// AddEdge will add a new edge between the provided vertices in the graph
func (g *UndirectedGraph) AddEdge(one, two int) {
	// Add vertices: one and two to the graph if they are not present
	g.AddVertex(one)
	g.AddVertex(two)

	// and finally add the edges: one->two and two->one for undirected graph
	g.edges[one][two] = struct{}{}
	g.edges[two][one] = struct{}{}
}
