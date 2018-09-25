package backtracking

import "fmt"

// Graph | Graph structure
type Graph struct {
	Vertices  int
	MaxColors int
	Adjacent  [][]int
	Color     []int
}

// NewGraph returns a new graph with corresponding details
func NewGraph(vertices, maxColors int) *Graph {
	g := &Graph{
		Vertices:  vertices,
		MaxColors: maxColors,
		Adjacent:  make([][]int, vertices),
		Color:     make([]int, vertices),
	}
	for i := range g.Adjacent {
		g.Color[i] = -1
		g.Adjacent[i] = make([]int, 0, vertices)
	}
	return g
}

// PrintColor prints the color of each vertices in the graph
func (g *Graph) PrintColor() {
	for i, node := range g.Color {
		fmt.Printf("Vertex %v has color %v\n", i, node)
	}
}

// AddEdge adds a edge from vertices u to v in the graph
func (g *Graph) AddEdge(u, v int) {
	g.Adjacent[u] = append(g.Adjacent[u], v)
	g.Adjacent[v] = append(g.Adjacent[v], u)
}

// IsSafe checks the adjacent vertices of u
// if any of the adjacent vertices has color `color`
// then the corresponding color is not safe for vertex u
func (g *Graph) IsSafe(u, color int) bool {
	for _, c := range g.Adjacent[u] {
		if g.Color[c] == color {
			return false
		}
	}
	return true
}

// SolveColoring uses backtracking to color the vertices
// using minimum different colors
func (g *Graph) SolveColoring(u int) bool {
	if u == g.Vertices {
		return true
	}
	for c := 1; c <= g.MaxColors; c++ {
		if g.IsSafe(u, c) {
			g.Color[u] = c
			if g.SolveColoring(u + 1) {
				return true
			}
			g.Color[u] = -1
		}
	}
	return false
}

// GraphColoring | run this in main to check the working
func GraphColoring() {
	graph := NewGraph(5, 3)
	/*
		Graph:
			3 - 0 - 1
		 	|	| /
			4	2
	*/
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(3, 4)

	if graph.SolveColoring(0) {
		fmt.Println("Graph Colored !!")
		graph.PrintColor()
	} else {
		fmt.Println("It is not possible to color the graph")
	}
	return
}
