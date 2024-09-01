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
)

type Vertex int

type Edge struct {
	Start  Vertex
	End    Vertex
	Weight int
}

func KruskalMST(n int, edges []Edge) ([]Edge, int) {
	// Initialize variables to store the minimum spanning tree and its total cost
	var mst []Edge
	var cost int

	// Create a new UnionFind data structure with 'n' nodes
	u := NewUnionFind(n)

	// Sort the edges in non-decreasing order based on their weights
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Iterate through the sorted edges
	for _, edge := range edges {
		// Check if adding the current edge forms a cycle or not
		if u.Find(int(edge.Start)) != u.Find(int(edge.End)) {
			// Add the edge to the minimum spanning tree
			mst = append(mst, edge)
			// Add the weight of the edge to the total cost
			cost += edge.Weight
			// Merge the sets containing the start and end vertices of the current edge
			u.Union(int(edge.Start), int(edge.End))
		}
	}

	// Return the minimum spanning tree and its total cost
	return mst, cost
}
