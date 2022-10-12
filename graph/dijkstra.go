package graph

import "github.com/TheAlgorithms/Go/sort"

type Item struct {
	node int
	dist int
}

func (a Item) More(b any) bool {
	// reverse direction for minheap
	return a.dist < b.(Item).dist
}
func (a Item) Idx() int {
	return a.node
}

func (g *Graph) Dijkstra(start, end int) (int, bool) {
	visited := make(map[int]bool)
	nodes := make(map[int]*Item)

	nodes[start] = &Item{
		dist: 0,
		node: start,
	}
	pq := sort.MaxHeap{}
	pq.Init(nil)
	pq.Push(*nodes[start])

	visit := func(curr Item) {
		visited[curr.node] = true
		for n, d := range g.edges[curr.node] {
			if visited[n] {
				continue
			}

			item := nodes[n]
			dist2 := curr.dist + d
			if item == nil {
				nodes[n] = &Item{node: n, dist: dist2}
				pq.Push(*nodes[n])
			} else if item.dist > dist2 {
				item.dist = dist2
				pq.Update(*item)
			}
		}
	}

	for pq.Size() > 0 {
		curr := pq.Pop().(Item)
		if curr.node == end {
			break
		}
		visit(curr)
	}

	item := nodes[end]
	if item == nil {
		return -1, false
	}
	return item.dist, true
}
