// KRUSKAL'S ALGORITHM

// Reference: Kruskal's Algorithm: https://www.scaler.com/topics/data-structures/kruskal-algorithm/

// Reference: Union Find Algorithm: https://www.scaler.com/topics/data-structures/disjoint-set/

// Author: Author: Mugdha Behere[https://github.com/MugdhaBehere]

// Worst Case Time Complexity: O(E log E), where E is the number of edges.

// Worst Case Space Complexity: O(V + E), where V is the number of vertices and E is the number of edges.

// see kruskal.go, kruskal_test.go

package graph

import (
	"sort"
	"fmt"
)

type Vertex int

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

func KruskalMST(n int, edges []Edge) ([]Edge, int) {
	// Initialize variables to store the minimum spanning tree and its total cost
	var mst []Edge
	var cost int

	// Create a new UnionFind data structure with 'n' nodes
	u := NewUnionFind(n)

	// Initialize each node in the UnionFind data structure
	for i := 0; i < n; i++ {
		u.parent[i] = Vertex(i)
		u.size[i] = 1
	}

	// Sort the edges in non-decreasing order based on their weights
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Iterate through the sorted edges
	for _, edge := range edges {
		// Check if adding the current edge forms a cycle or not
		if u.Find(edge.Start) != u.Find(edge.End) {
			// Add the edge to the minimum spanning tree
			mst = append(mst, edge)
			// Add the weight of the edge to the total cost
			cost += edge.Weight
			// Merge the sets containing the start and end vertices of the current edge
			u = u.Union(edge.Start, edge.End)
		}
	}

	// Return the minimum spanning tree and its total cost
	return mst, cost
}

// The main function sets up a sample graph and finds its minimum spanning tree using Kruskal's algorithm
func main() {
	// Define the number of vertices in the graph
	n := 5

	// Define the edges of the graph
	edges := []Edge{
		{Start: 0, End: 1, Weight: 4},
		{Start: 0, End: 2, Weight: 13},
		{Start: 0, End: 3, Weight: 7},
		{Start: 0, End: 4, Weight: 7},
		{Start: 1, End: 2, Weight: 9},
		{Start: 1, End: 3, Weight: 3},
		{Start: 1, End: 4, Weight: 7},
		{Start: 2, End: 3, Weight: 10},
		{Start: 2, End: 4, Weight: 14},
		{Start: 3, End: 4, Weight: 4},
	}

	// Find the minimum spanning tree and its cost using Kruskal's algorithm
	mst, cost := KruskalMST(n, edges)

	// Print the edges in the minimum spanning tree and the total cost
	fmt.Println("Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("Start: %d, End: %d, Weight: %d\n", edge.Start, edge.End, edge.Weight)
	}
	fmt.Println("Total cost:", cost)
}