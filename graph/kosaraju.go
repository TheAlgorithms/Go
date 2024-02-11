// Kosaraju Algorithm is used to get strongly connected components of a graph
// Time Complexity: O(N + M), where N is number of vertices and M is number of edges
// Space Complexity: O(N)
// Reference: https://www.baeldung.com/cs/kosaraju-algorithm-scc

package graph

func topologicalSort(node int, adj [][]int, sortedNodes *[]int, visited map[int]bool) {
	visited[node] = true
	for _, nbr := range adj[node] {
		if visited[nbr] {
			continue
		}
		topologicalSort(nbr, adj, sortedNodes, visited)
	}
	*sortedNodes = append(*sortedNodes, node)
}

func getTransposeGraph(n int, adj [][]int) [][]int {
	transposeGraph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for _, nbr := range adj[i] {
			transposeGraph[nbr] = append(transposeGraph[nbr], i)
		}
	}
	return transposeGraph
}

func dfs(node int, adj [][]int, visited map[int]bool, currentSCC *[]int) {
	visited[node] = true
	*currentSCC = append(*currentSCC, node)
	for _, nbr := range adj[node] {
		if visited[nbr] {
			continue
		}
		dfs(nbr, adj, visited, currentSCC)
	}
}

// Get Strongly Connected Components using Kosaraju Algorithm
func GetStronglyConnectedComponents(n int, adj [][]int) [][]int {
	transposeGraph := getTransposeGraph(n, adj)
	sortedNodes := []int{}
	visited := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		topologicalSort(i, adj, &sortedNodes, visited)
	}
	visited = make(map[int]bool)
	allSCC := [][]int{}
	for i := len(sortedNodes) - 1; i >= 0; i-- {
		currentSCC := []int{}
		if visited[sortedNodes[i]] {
			continue
		}
		dfs(sortedNodes[i], transposeGraph, visited, &currentSCC)
		allSCC = append(allSCC, currentSCC)
	}
	return allSCC
}
