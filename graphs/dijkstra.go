package graph

// Dijkstra returns the shortest path from startNode to all the other Nodes
func Dijkstra(g *Graph, startNode int) map[int]int {
	costs, visited := InitState(g.Nodes, startNode)

	for i := 0; i < len(g.Nodes); i++ {

		// Get closest non visited Node
		node := getClosestNode(costs, visited)

		visited[node] = true

		nodeEdges := g.GetEdgesFrom(node)

		for _, edge := range nodeEdges {

			newCostForNeighbor := costs[node] + edge.Cost

			// if the distance above is less than the distance currently in the costs for the neighbor
			if newCostForNeighbor < costs[edge.To] {
				// Update the costs for that neighbor
				costs[edge.To] = newCostForNeighbor
			}
		}
	}

	return costs
}

// InitState returns an initialized cost and visited tables for the Dijkstra algorithm to work with
func InitState(nodes []int, startNode int) (map[int]int, map[int]bool) {
	costs := make(map[int]int)
	visited := make(map[int]bool)

	for _, node := range nodes {
		if node != startNode {
			costs[node] = int(^uint(0) >> 1)
		} else {
			costs[node] = 0
		}
		visited[node] = false
	}

	return costs, visited
}

// getClosestNode returns the closest Node from the costs table that has not been visited yet
func getClosestNode(costs map[int]int, visited map[int]bool) int {
	minCost, closestNode := int(^uint(0)>>1), -1

	for node, cost := range costs {
		if !visited[node] && minCost >= cost {
			minCost = cost
			closestNode = node
		}
	}

	return closestNode
}
