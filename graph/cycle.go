package graph

func (g *Graph) HasCycle() bool {

	all := map[int]struct{}{}
	visiting := map[int]struct{}{}
	visited := map[int]struct{}{}

	for v := range g.edges {
		all[v] = struct{}{}
	}

	for current := range all {
		if g.HasCycleHelper(current, all, visiting, visited) {
			return true
		}
	}

	return false

}

func (g Graph) HasCycleHelper(v int, all, visiting, visited map[int]struct{}) bool {
	delete(all, v)
	visiting[v] = struct{}{}

	neighbors := g.edges[v]
	for v := range neighbors {
		if _, ok := visited[v]; ok {
			continue
		} else if _, ok := visiting[v]; ok {
			return true
		} else if g.HasCycleHelper(v, all, visiting, visited) {
			return true
		}
	}
	delete(visiting, v)
	visited[v] = struct{}{}
	return false
}
