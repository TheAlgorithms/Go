package breathfirstsearch

/*

	Breadth First Search
	https://en.wikipedia.org/wiki/Breadth-first_search

*/

func BreadthFirstSearch(start, end, nodes int, edges [][]int) (isConnected bool, distance int) {
	queue := make([]int, 0)
	discovered := make([]int, nodes)
	discovered[start] = 1
	queue = append(queue, start)
	for len(queue) > 0 {
		v := queue[0]
		if len(queue) > 0 {
			queue = queue[1:]
		}
		for i := 0; i < len(edges[v]); i++ {
			if discovered[i] == 0 && edges[v][i] > 0 {
				if i == end {
					return true, discovered[v]
				}
				discovered[i] = discovered[v] + 1
				queue = append(queue, i)
			}
		}
	}
	return false, 0
}
