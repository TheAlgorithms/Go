package graph

import "strconv"

type Edge struct {
	From int
	To   int
	Cost int
}

type Graph struct {
	Edges []*Edge
	Nodes []int
}

// InsertEdge adds an Edge to the Graph
func (g *Graph) InsertEdge(from int, to int, cost int) {
	edge := Edge{
		From: from,
		To:   to,
		Cost: cost,
	}
	g.InsertNode(from)
	g.InsertNode(to)
	g.Edges = append(g.Edges, &edge)
}

// InsertNode adds a Node to the Graph list of Nodes, if the the node wasn't already added
func (g *Graph) InsertNode(node int) {
	if !g.ContainsNode(node) {
		g.Nodes = append(g.Nodes, node)
	}
}

// ContainsNode checks if graph already contains a node with given index
func (g *Graph) ContainsNode(targetNode int) bool {
	for _, node := range g.Nodes {
		if node == targetNode {
			return true
		}
	}
	return false
}

// GetEdgesFrom returns all the Edges that start with the specified Node
func (g *Graph) GetEdgesFrom(node int) []*Edge {
	var edges []*Edge
	for _, edge := range g.Edges {
		if edge.From == node {
			edges = append(edges, edge)
		}
	}
	return edges
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var graphStr string

	graphStr += "Edges:\n"
	for _, edge := range g.Edges {
		graphStr += strconv.Itoa(edge.From) + " -> " + strconv.Itoa(edge.To) + " = " + strconv.Itoa(edge.Cost)
		graphStr += "\n"
	}
	graphStr += "\n"

	graphStr += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			graphStr += strconv.Itoa(node)
		} else {
			graphStr += strconv.Itoa(node) + ", "
		}
	}
	graphStr += "\n"

	return graphStr
}
