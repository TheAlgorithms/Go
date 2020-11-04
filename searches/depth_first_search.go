package searches

import "fmt"

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

func dfs(start, end int, nodes []int, edges [][]bool) ([]int, bool) {
	var route []int
	var stack []int
	startIdx := getIdx(start, nodes)
	stack = append(stack, startIdx)
	for len(stack) > 0 {
		now := stack[len(stack)-1]
		route = append(route, nodes[now])
		if len(stack) > 1 {
			stack = stack[:len(stack)-1]
		} else {
			stack = stack[:len(stack)-1]
		}
		for i := 0; i < len(edges[now]); i++ {
			if edges[now][i] && notExist(i, stack) {
				stack = append(stack, i)
			}
			edges[now][i] = false
			edges[i][now] = false
		}
		if route[len(route)-1] == end {
			return route, true
		}
	}
	return nil, false
}

func main() {
	nodes := []int{
		1, 2, 3, 4, 5, 6,
	}
	/*
		sample graph
		①-②
		|  |
		③-④-⑤-⑥
	*/
	edges := [][]bool{
		{false, true, true, false, false, false},
		{true, false, false, true, false, false},
		{true, false, false, true, false, false},
		{false, true, true, false, true, false},
		{false, false, false, true, false, true},
		{false, false, false, false, true, false},
	}
	start := 1
	end := 6
	route, _ := dfs(start, end, nodes, edges)
	fmt.Println(route)
}
