// KRUSKAL'S ALGORITHM
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// https://cp-algorithms.com/graph/mst_kruskal_with_dsu.html

package graph

import (
	"sort"
)

type Vertex int

// Edge describes the edge of a weighted graph
type Edge struct {
	Start  Vertex
	End    Vertex
	Weight int
}

// DisjointSetUnionElement describes what an element of DSU looks like
type DisjointSetUnionElement struct {
	Parent Vertex
	Rank   int
}

// DisjointSetUnion is a data structure that treats its elements as separate sets
// and provides fast operations for set creation, merging sets, and finding the parent
// of the given element of a set.
type DisjointSetUnion []DisjointSetUnionElement

// NewDSU will return an initialised DSU using the value of n
// which will be treated as the number of elements out of which
// the DSU is being made
func NewDSU(n int) *DisjointSetUnion {

	dsu := DisjointSetUnion(make([]DisjointSetUnionElement, n))
	return &dsu
}

// MakeSet will create a set in the DSU for the given node
func (dsu DisjointSetUnion) MakeSet(node Vertex) {

	dsu[node].Parent = node
	dsu[node].Rank = 0
}

// FindSetRepresentative will return the parent element of the set the given node
// belongs to. Since every single element in the path from node to parent
// has the same parent, we store the parent value for each element in the
// path. This reduces consequent function calls and helps in going from O(n)
// to O(log n). This is known as path compression technique.
func (dsu DisjointSetUnion) FindSetRepresentative(node Vertex) Vertex {

	if node == dsu[node].Parent {
		return node
	}

	dsu[node].Parent = dsu.FindSetRepresentative(dsu[node].Parent)
	return dsu[node].Parent
}

// unionSets will merge two given sets. The naive implementation of this
// always combines the secondNode's tree with the firstNode's tree. This can lead
// to creation of trees of length O(n) so we optimize by attaching the node with
// smaller rank to the node with bigger rank. Rank represents the upper bound depth of the tree.
func (dsu DisjointSetUnion) UnionSets(firstNode Vertex, secondNode Vertex) {

	firstNode = dsu.FindSetRepresentative(firstNode)
	secondNode = dsu.FindSetRepresentative(secondNode)

	if firstNode != secondNode {

		if dsu[firstNode].Rank < dsu[secondNode].Rank {
			firstNode, secondNode = secondNode, firstNode
		}
		dsu[secondNode].Parent = firstNode

		if dsu[firstNode].Rank == dsu[secondNode].Rank {
			dsu[firstNode].Rank++
		}
	}
}

// KruskalMST will return a minimum spanning tree along with its total cost
// to using Kruskal's algorithm. Time complexity is O(m * log (n)) where m is
// the number of edges in the graph and n is number of nodes in it.
func KruskalMST(n int, edges []Edge) ([]Edge, int) {

	var mst []Edge // The resultant minimum spanning tree
	var cost int = 0

	dsu := NewDSU(n)

	for i := 0; i < n; i++ {
		dsu.MakeSet(Vertex(i))
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	for _, edge := range edges {

		if dsu.FindSetRepresentative(edge.Start) != dsu.FindSetRepresentative(edge.End) {

			mst = append(mst, edge)
			cost += edge.Weight
			dsu.UnionSets(edge.Start, edge.End)
		}
	}

	return mst, cost
}
