// This file contains the simple structural implementation of
// directed & undirected graphs used within the graph package
// Author(s): [Shivam](https://github.com/Shivam010), [Tahmeed](https://github.com/Tahmeed156)

package graph

// Graph provides a structure to store the graph.
// It is safe to use its empty object.
type Graph struct {
	vertices int
	edges    map[int]map[int]int // Stores weight of an edge
	directed bool                // Differentiate directed/undirected graphs
}

// Constructor functions for graphs
func MakeUndirectedGraph(v int) Graph {
	return Graph{
		directed: false,
		vertices: v,
	}
}
func MakeDirectedGraph(v int) Graph {
	return Graph{
		directed: true,
		vertices: v,
	}
}

// AddVertex will add a new vertex in the graph.
// If the vertex already exists it will do nothing.
func (g *Graph) AddVertex(v int) {
	if g.edges == nil {
		g.edges = make(map[int]map[int]int)
	}

	// Check if vertex is present or not
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[int]int)
	}
}

// AddEdge will add a new edge between the provided vertices in the graph
func (g *Graph) AddEdge(one, two int) {
	// Add an edge with 0 weight
	g.AddWeightedEdge(one, two, 0)
}

// AddWeightedEdge will add a new weighted edge between the provided vertices in the graph
func (g *Graph) AddWeightedEdge(one, two, weight int) {
	// Add vertices: one and two to the graph if they are not present
	g.AddVertex(one)
	g.AddVertex(two)

	// And finally add the edges
	// one->two and two->one for undirected graph
	// one->two for directed graphs
	g.edges[one][two] = weight
	if !g.directed {
		g.edges[two][one] = weight
	}
}
