// This file contains the graph coloring implementation using BFS
// time complexity: O(V+E) where V is the number of vertices and E is the number of edges in the graph
// space complexity: O(V) where V is the number of vertices in the graph
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring

import "container/list"

// ColorUsingBFS will return the Color of each vertex and the
// total number of different colors used, using BFS
func (g *Graph) ColorUsingBFS() (map[int]Color, int) {
	// Initially all vertices will have same color
	vertexColors := make(map[int]Color, g.vertices)
	for i := 0; i < g.vertices; i++ {
		vertexColors[i] = 1
	}

	visited := make(map[int]struct{})
	// Run BFS from each non-visited vertex
	for i := 0; i < g.vertices; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		visited[i] = struct{}{}

		queue := list.New()
		queue.PushBack(i)

		for queue.Len() != 0 {
			// front vertex in the queue
			frontNode := queue.Front()
			front := frontNode.Value.(int)
			queue.Remove(frontNode)

			// Now, check all neighbours of front vertex, if they have same
			// color as that of front, change their color
			for nb := range g.edges[front] {
				if vertexColors[nb] == vertexColors[front] {
					vertexColors[nb]++
				}
				// if the neighbour is not already visited, add it to the queue
				if _, ok := visited[nb]; !ok {
					visited[nb] = struct{}{}
					queue.PushBack(nb)
				}
			}
		}
	}

	colorsUsed := 0
	for _, cr := range vertexColors {
		if colorsUsed < int(cr) {
			colorsUsed = int(cr)
		}
	}
	return vertexColors, colorsUsed
}
