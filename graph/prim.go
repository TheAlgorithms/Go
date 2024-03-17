// The Prim's algorithm computes the minimum spanning tree for a weighted undirected graph
// Worst Case Time Complexity: O(E log V) using Binary heap, where V is the number of vertices and E is the number of edges
// Space Complexity: O(V + E)
// Implementation is based on the book 'Introduction to Algorithms' (CLRS)

package graph

import (
	"container/heap"
)

type minEdge []Edge

func (h minEdge) Len() int           { return len(h) }
func (h minEdge) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h minEdge) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minEdge) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *minEdge) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (g *Graph) PrimMST(start Vertex) ([]Edge, int) {
	var mst []Edge
	marked := make([]bool, g.vertices)
	h := &minEdge{}
	// Pushing neighbors of the start node to the binary heap
	for neighbor, weight := range g.edges[int(start)] {
		heap.Push(h, Edge{start, Vertex(neighbor), weight})
	}
	marked[start] = true
	cost := 0
	for h.Len() > 0 {
		e := heap.Pop(h).(Edge)
		end := int(e.End)
		// To avoid cycles
		if marked[end] {
			continue
		}
		marked[end] = true
		cost += e.Weight
		mst = append(mst, e)
		// Check for neighbors of the newly added edge's End vertex
		for neighbor, weight := range g.edges[end] {
			if !marked[neighbor] {
				heap.Push(h, Edge{e.End, Vertex(neighbor), weight})
			}
		}
	}
	return mst, cost
}
