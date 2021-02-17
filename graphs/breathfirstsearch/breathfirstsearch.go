package breathfirstsearch

func getIdx(target int, nodes []int) int {
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == target {
			return i
		}
	}
	return -1
}

func notExist(target int, slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return false
		}
	}
	return true
}

func breadthFirstSearch(start, end int, nodes []int, edges [][]bool) bool {
	var route []int
	var queue []int
	startIdx := getIdx(start, nodes)
	queue = append(queue, startIdx)
	for len(queue) > 0 {
		now := queue[0]
		route = append(route, nodes[now])
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = queue[0:]
		}
		for i := 0; i < len(edges[now]); i++ {
			if edges[now][i] && notExist(i, queue) {
				queue = append(queue, i)
			}
			edges[now][i] = false
			edges[i][now] = false
		}
		if route[len(route)-1] == end {
			return true
		}
	}
	return false
}

// func main() {
// 	nodes := []int{
// 		1, 2, 3, 4, 5, 6,
// 	}
// 	/*
// 		sample graph
// 		①-②
// 		|  |
// 		③-④-⑤-⑥
// 	*/
// 	edges := [][]bool{
// 		{false, true, true, false, false, false},
// 		{true, false, false, true, false, false},
// 		{true, false, false, true, false, false},
// 		{false, true, true, false, true, false},
// 		{false, false, false, true, false, true},
// 		{false, false, false, false, true, false},
// 	}
// 	start := 1
// 	end := 6
// 	result := breadthFirstSearch(start, end, nodes, edges)
// 	fmt.Println(result)
// }
