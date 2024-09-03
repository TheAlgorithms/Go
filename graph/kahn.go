package graph

func Kahn(n int, dependencies [][]int) []int {
	g := Graph{vertices: n, Directed: true}
	inDegree := make([]int, n)

	for _, d := range dependencies {
		if _, ok := g.edges[d[0]][d[1]]; !ok {
			g.AddEdge(d[0], d[1])
			inDegree[d[1]]++
		}
	}

	queue := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, n)

	for len(queue) > 0 {
		vtx := queue[0]
		queue = queue[1:]
		order = append(order, vtx)
		for neighbour := range g.edges[vtx] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	if len(order) != n {
		return nil
	}
	return order
}
