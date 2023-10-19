// Union Find Algorithm or Dynamic Connectivity algorithm, often implemented with the help 

//of the union find data structure,

// is used to efficiently maintain connected components in a graph that undergoes dynamic changes,

// such as edges being added or removed over time

// Worst Case Time Complexity: The time complexity of find operation is nearly constant or 

//O(α(n)), where where α(n) is the inverse Ackermann function

// practically, this is a very slowly growing function making the time complexity for find 

//operation nearly constant.

// The time complexity of the union operation is also nearly constant or O(α(n))

// Worst Case Space Complexity: O(n), where n is the number of nodes or element in the structure

// Reference: https://www.scaler.com/topics/data-structures/disjoint-set/

// Author: Mugdha Behere[https://github.com/MugdhaBehere]

// see: unionfind.go, unionfind_test.go

package graph

import "fmt"

// Defining the union-find data structure
type UnionFind struct {
	parent []int
	size []int 
}


// Initialise a new union find data structure with s nodes
func NewUnionFind(s int) UnionFind {
	parent := make([]int, s)
	size := make([]int, s)
	for k := 0; k < s; k++ {
		parent[k] = k
		size[k] = 1
	}
	return UnionFind{parent, size}
}


// to find the root of the set to which the given element belongs, the Find function serves the purpose
func (u UnionFind) Find(q int) int {
	for q != u.parent[q] {
		q = u.parent[q]
	}
	return q
}


// to merge two sets to which the given elements belong, the Union function serves the purpose
func (u UnionFind) Union(a, b int) UnionFind {
	rootP := u.Find(a)
	rootQ := u.Find(b)

	if rootP == rootQ {
		return u
	}

	if u.size[rootP] < u.size[rootQ] {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	} else {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	}
	return u
}


