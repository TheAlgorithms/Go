// Package search Graph search algorithms
// reference: https://en.wikipedia.org/wiki/Tree_traversal
package search

import (
	"github.com/TheAlgorithms/Go/datastructures/stack"
)

/*DepthFirstSearch is an algorithm for traversing or searching graph data structures.
	The algorithm starts at the start node and explores as far as possible along each branch before backtracking.

	Worst-case performance			O(|V|+|E|)
	Worst-case space complexity		O(|V|)
		- |V| the number of vertices in graph
		- |E| the number of edges in graph
reference: https://en.wikipedia.org/wiki/Depth-first_search
*/
func DepthFirstSearch(start, end, nodes int, edges [][]int) (isConnected bool, distance int) {
	discovered := make([]int, nodes)
	discovered[start] = 1
	s := new(stack.Stack)
	s.Push(start)
	for !s.IsEmpty() {
		v, isOK := s.Pop()
		if !isOK {
			return false, 0
		}
		for i := 0; i < len(edges[v.(int)]); i++ {
			if discovered[i] == 0 && edges[v.(int)][i] > 0 {
				if i == end {
					return true, discovered[v.(int)]
				}
				discovered[i] = discovered[v.(int)] + 1
				s.Push(i)
			}
		}
	}
	return false, 0
}
