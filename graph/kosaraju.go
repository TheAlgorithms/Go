package graph

import (
	"github.com/abaron10/Gothon/gothonSlice"
)

type graph struct {
	nodes    []int
	reversed *graph
	edges    map[int][]*edge
}

type edge struct {
	from int
	to   int
}

func newEdge(from int, to int) *edge {
	return &edge{from: from, to: to}
}

func NewGraphK() *graph {
	reversedGraph := &graph{nodes: []int{}, edges: map[int][]*edge{}}
	return &graph{nodes: []int{}, edges: map[int][]*edge{}, reversed: reversedGraph}
}

func (g *graph) AddEdge(from int, to int) {
	edges, _ := g.edges[from]
	edges = append(edges, newEdge(from, to))

	g.edges[from] = edges
	g.nodes = append(g.nodes, from)

	//Computing transposed graph at the same time user adds an Edge.
	if g.reversed != nil {
		g.reversed.AddEdge(to, from)
	}
}

func (g *graph) EvaluateSCC() [][]int {
	baseOrder := g.calculateStackBaseOrder()
	sccResult := g.findSCCComponents(baseOrder)
	return g.computeAnswer(sccResult)
}

func (g *graph) calculateStackBaseOrder() []int {
	visited := map[int]struct{}{}
	stack := []int{}

	for _, node := range g.nodes {
		if _, ok := visited[node]; !ok {
			g.dfs(node, visited, &stack)
		}
	}

	return stack
}

func (g *graph) findSCCComponents(orderNodes []int) map[int]int {
	visited := map[int]struct{}{}
	scc := map[int]int{}
	counter := 0

	for len(orderNodes) > 0 {
		//using https://github.com/abaron10/Gothon library to emulate stack popping as Python do
		node := gothonSlice.Pop(&orderNodes, -1)
		if _, ok := visited[node]; !ok {
			g.dfsReversed(node, visited, scc, counter)
			counter++
		}
	}

	return scc
}

/*
---------------------------------------------------------------------------------------------------------
--------------------------------DFS IMPLEMENTATIONS (BASE AND TRANSPOSED GRAPH)---------------------------
-----------------------------------------------------------------------------------------------------------
*/
func (g *graph) dfs(from int, visited map[int]struct{}, stack *[]int) {
	visited[from] = struct{}{}
	edges, ok := g.edges[from]
	if ok {
		for _, edge := range edges {
			if _, ok = visited[edge.to]; !ok {
				g.dfs(edge.to, visited, stack)
			}
		}
	}
	*stack = append(*stack, from)
}

func (g *graph) dfsReversed(from int, visited map[int]struct{}, scc map[int]int, counter int) {
	visited[from] = struct{}{}
	scc[from] = counter
	edges, ok := g.reversed.edges[from]
	if ok {
		for _, edge := range edges {
			if _, ok = visited[edge.to]; !ok {
				g.dfsReversed(edge.to, visited, scc, counter)
			}
		}
	}
	scc[from] = counter
}

/*---------------------------------------------------------------------------------------------------------
--------------------------------HELPER METHOD TO COMPUTE A READABLE SCC ANSWER-----------------------------
-----------------------------------------------------------------------------------------------------------
*/

func (g *graph) PopulateGraph(graph map[int][]int) {
	for node, edges := range graph {
		for _, edge := range edges {
			g.AddEdge(node, edge)
		}
	}
}

func (g *graph) computeAnswer(sccKeys map[int]int) [][]int {
	response := map[int][]int{}
	for node, sscComponent := range sccKeys {
		response[sscComponent] = append(response[sscComponent], node)
	}

	scc := [][]int{}
	for _, value := range response {
		scc = append(scc, value)
	}
	return scc
}
