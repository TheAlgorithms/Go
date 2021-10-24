package graph

import "container/heap"

type Item struct {
	node  int
	dist  int
	index int
}

type Pqueue []*Item

func (p Pqueue) Len() int { return len(p) }
func (p Pqueue) Less(i, j int) bool {
	if p[i].dist == p[j].dist {
		return p[i].node < p[j].node
	}
	return p[i].dist < p[j].dist
}
func (p Pqueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = j
	p[j].index = i
}
func (p *Pqueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}
func (p *Pqueue) Pop() interface{} {
	old := *p
	n := len(old)

	item := old[n-1]
	old[n-1] = nil
	item.index = -1

	*p = old[0 : n-1]
	return item
}
func (p *Pqueue) update(item *Item, dist int) {
	item.dist = dist
	heap.Fix(p, item.index-1)
}

func (g *UndirectedGraph) Dijkstra(start, end int) (int, bool) {
	visited := make(map[int]bool)
	nodes := make(map[int]*Item)

	pq := make(Pqueue, 0, len(g.edges))
	nodes[start] = &Item{
		dist:  0,
		index: 0,
		node:  start,
	}
	pq = append(pq, nodes[start])
	heap.Init(&pq)

	for len(pq) > 0 {
		curr := *heap.Pop(&pq).(*Item)
		visited[curr.node] = true

		if curr.node == end {
			break
		}

		for n, d := range g.edges[curr.node] {
			if visited[n] {
				continue
			}

			item := nodes[n]
			dist2 := curr.dist + d
			if item == nil {
				nodes[n] = &Item{node: n, dist: dist2}
				heap.Push(&pq, nodes[n])
			} else if item.dist > dist2 {
				pq.update(item, dist2)
			}
		}
	}

	item := nodes[end]
	if item == nil {
		return -1, false
	} else {
		return item.dist, true
	}
}
