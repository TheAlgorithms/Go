// Union Find Algorithm or Dynamic Connectivity algorithm, often implemented with the help
//of the union find data structure,
// is used to efficiently maintain connected components in a graph that undergoes dynamic changes,
// such as edges being added or removed over time
// Worst Case Time Complexity: The time complexity of find operation is nearly constant or
//O(α(n)), where α(n) is the inverse Ackermann function
// practically, this is a very slowly growing function making the time complexity for find
//operation nearly constant.
// The time complexity of the union operation is also nearly constant or O(α(n))
// Worst Case Space Complexity: O(n), where n is the number of nodes or element in the structure
// Reference: https://www.scaler.com/topics/data-structures/disjoint-set/
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
// Author: Mugdha Behere[https://github.com/MugdhaBehere]
// see: unionfind.go, unionfind_test.go

package graph

// Defining the union-find data structure
type UnionFind struct {
	parent []int
	rank   []int
}

// Initialise a new union find data structure with s nodes
func NewUnionFind(s int) UnionFind {
	parent := make([]int, s)
	rank := make([]int, s)
	for i := 0; i < s; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return UnionFind{parent, rank}
}

// Find finds the root of the set to which the given element belongs.
// It performs path compression to make future Find operations faster.
func (u *UnionFind) Find(q int) int {
	if q != u.parent[q] {
		u.parent[q] = u.Find(u.parent[q])
	}
	return u.parent[q]
}

// Union merges the sets, if not already merged, to which the given elements belong.
// It performs union by rank to keep the tree as flat as possible.
func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	if u.rank[rootP] < u.rank[rootQ] {
		u.parent[rootP] = rootQ
	} else if u.rank[rootP] > u.rank[rootQ] {
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootQ] = rootP
		u.rank[rootP]++
	}
}
