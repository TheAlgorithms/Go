// Package disjointsets
// It provides an implementation of Disjoint Set Union (DSU) data structure,
// which allows you to efficiently manages disjoint sets, enabling operations like finding set representatives and merging sets,
// commonly used in graph algorithms and dynamic connectivity problems.

// To learn more about Disjoint Set Uniong (DSU) follow the links below:
// cp-algorithms.com blog : https://cp-algorithms.com/data_structures/disjoint_set_union.html

// To practice problems on DSU follow the links below:
// USACO Guide: https://usaco.guide/gold/dsu?lang=cpp

// authors [Gagan Gaurav](https://github.com/gagan-gaurav)

package disjointsets

// Dsu represents the Disjoint Set Union data structure.
type Dsu struct {
	e []int // The parent/size array to represent disjoint sets
}

// NewDisjointSet creates a new DSU instance with the given size and initializes it.
func NewDisjointSet(size int) *Dsu {
	// Initialize the parent/size array with all elements as -1, indicating that each element is its own set.
	arr := make([]int, size)
	for i := range arr {
		arr[i] = -1
	}
	dsu := Dsu{
		e: arr,
	}
	return &dsu
}

// Get returns the representative element (root) of the set containing element x.
// It uses path compression for efficiency.
func (dsu *Dsu) Get(x int) int {
	if dsu.e[x] < 0 {
		return x
	}
	// Path compression: Update the parent of x to the representative element of its set.
	dsu.e[x] = dsu.Get(dsu.e[x])
	return dsu.e[x]
}

// IsInSameSet checks if elements a and b belong to the same set.
func (dsu *Dsu) IsInSameSet(a int, b int) bool {
	return dsu.Get(a) == dsu.Get(b)
}

// Size returns the size (number of elements) of the set containing element x.
func (dsu *Dsu) Size(x int) int {
	// The size of a set is stored as a negative value in the parent/size array.
	return -dsu.e[dsu.Get(x)]
}

// Union performs the union of two sets containing elements a and b using the "union by size" heuristic.
// It returns true if the sets were successfully merged, and false if a and b were already in the same set.
func (dsu *Dsu) Union(a int, b int) bool {
	x := dsu.Get(a)
	y := dsu.Get(b)
	if x == y {
		// Elements a and b are already in the same set.
		return false
	}
	if dsu.e[x] > dsu.e[y] {
		// Attach the smaller set (y) to the larger set (x).
		dsu.e[x], dsu.e[y] = dsu.e[y], dsu.e[x]
	}
	// Update the size of the set and make x the parent of y.
	dsu.e[x] += dsu.e[y]
	dsu.e[y] = x
	return true
}
