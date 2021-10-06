package graph

import (
	"sort"
)

// edge describes the edge of a weighted graph
type edge struct {
	start  int
	end    int
	weight int
}

// disjointSetUnion is a data structure that treats its elements as separate sets
// and provides fast operations for set creation, merging sets, and finding the parent
// of the given element of a set.
type disjointSetUnion struct {
	parent []int
	rank   []int
}

// newDSU will return an initialised DSU using the value of n
// which will be treated as the number of elements out of which
// the DSU is being made
func newDSU(n int) *disjointSetUnion {

	return &disjointSetUnion{

		parent: make([]int, n),
		rank:   make([]int, n),
	}
}

// makeSet will create a set in the DSU for the given node
func (dsu *disjointSetUnion) makeSet(node int) {

	dsu.parent[node] = node
	dsu.rank[node] = 0
}

// findSet will return the parent element of the set the given node
// belongs to. Since every single element in the path from node to parent
// has the same parent, we store the parent value for each element in the
// path. This reduces consequent function calls and helps in going from O(n)
// to O(log n). This is known as path compression technique.
func (dsu *disjointSetUnion) findSet(node int) int {

	if node == dsu.parent[node] {
		return node
	}

	dsu.parent[node] = dsu.findSet(dsu.parent[node])
	return dsu.parent[node]
}

// unionSets will merge two given sets. The naive implementation of this
// always combines the secondNode's tree with the firstNode's tree. This can lead
// to creation of trees of length O(n) so we optimize by attaching the node with
// smaller rank to the node with bigger rank. Rank represents the upper bound depth of the tree.
func (dsu *disjointSetUnion) unionSets(firstNode int, secondNode int) {

	firstNode = dsu.findSet(firstNode)
	secondNode = dsu.findSet(secondNode)

	if firstNode != secondNode {

		if dsu.rank[firstNode] < dsu.rank[secondNode] {
			swap(firstNode, secondNode)
		}
		dsu.parent[secondNode] = firstNode

		if dsu.rank[firstNode] == dsu.rank[secondNode] {
			dsu.rank[firstNode]++
		}
	}
}

// KruskalMST will return a minimum spanning tree along with its total cost
// to using Kruskal's algorithm. Time complexity is O(m * log (n)) where m is
// the number of edges in the graph and n is number of nodes in it.
func KruskalMST(n int, edges []edge) ([]edge, int) {

	var mst []edge // The resultant minimum spanning tree
	var cost int = 0

	dsu := newDSU(n)

	for i := 0; i < n; i++ {
		dsu.makeSet(i)
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	for _, edge := range edges {

		if dsu.findSet(edge.start) != dsu.findSet(edge.end) {

			mst = append(mst, edge)
			cost += edge.weight
			dsu.unionSets(edge.start, edge.end)
		}
	}

	return mst, cost
}

// swap is a utility function to swap two numbers
func swap(a int, b int) {

	temp := a
	a = b
	b = temp
}
