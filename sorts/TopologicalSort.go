//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"fmt"
)

type Dependency struct {
	nodeId   int
	children []int
}

// converts dependencies to edges
func dependenciesToEdges(dependencies []Dependency) []Edge {
	edges := []Edge{}
	for _, node := range dependencies {
		for _, child := range node.children {
			edges = append(edges, Edge{start: node.nodeId, end: child})
		}
	}
	return edges
}

type Edge struct {
	start int // start and end are node ids
	end   int
}

type Node struct {
	id       int
	children []int
	selected bool
	incoming int
}

// takes a list of edges as input with start and end node_ids
// useful for cases like task ordering
func topologicalSort(input []Edge) ([]int, [][]int) {

	// generate a mapping for node_id and number of incoming edges
	incoming := make(map[int]int)
	for _, edge := range input {
		if _, ok := incoming[edge.start]; !ok {
			incoming[edge.start] = 0
		}
		incoming[edge.end] += 1
	}

	nodes := make([]*Node, 0)
	for k, v := range incoming {
		chdn := []int{}
		for _, edge := range input {
			if edge.start == k {
				chdn = append(chdn, edge.end)
			}
		}
		nodes = append(nodes, &Node{id: k, children: chdn, selected: false, incoming: v})
	}

	res := []int{}
	var levels [][]int
	for len(res) < len(nodes) {
		level := []int{}
		queue := []int{}
		for idx, node := range nodes {
			// select a node with no incoming dependency
			if !node.selected && node.incoming == 0 {
				queue = append(queue, idx)
				nodes[idx].selected = true
			}
		}
		for _, idx := range queue {
			node := nodes[idx]
			level = append(level, node.id)
			// decrement incoming values for all children
			for _, childId := range node.children {
				for _, child := range nodes {
					if !child.selected && child.id == childId {
						child.incoming--
					}
				}
			}
		}
		res = append(res, level...)
		levels = append(levels, level)
	}
	return res, levels
}

func main() {
	var edges []Edge
	edges = append(edges, Edge{1, 2})
	edges = append(edges, Edge{2, 3})
	edges = append(edges, Edge{4, 3})
	edges = append(edges, Edge{4, 5})
	edges = append(edges, Edge{3, 6})
	edges = append(edges, Edge{5, 6})

	order, levels := topologicalSort(edges)
	fmt.Println("result: ", order)
	fmt.Println("levels: ", levels)

	// or you can specify dependencies
	var deps []Dependency
	deps = append(deps, Dependency{nodeId: 1, children: []int{2}})
	deps = append(deps, Dependency{nodeId: 2, children: []int{3}})
	deps = append(deps, Dependency{nodeId: 4, children: []int{3, 5}})
	deps = append(deps, Dependency{nodeId: 5, children: []int{6}})
	deps = append(deps, Dependency{nodeId: 3, children: []int{6}})

	edges = dependenciesToEdges(deps)
	order, levels = topologicalSort(edges)
	fmt.Println("result: ", order)
	fmt.Println("levels: ", levels)

}
