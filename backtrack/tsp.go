// tsp.go
// description: solution of the traveling salesman problem
// details:
// Given a list of cities and the distances between each pair of cities,
// what is the shortest possible route that
// visits each city exactly once and returns to the origin city?
// wikipedia: https://en.wikipedia.org/wiki/Travelling_salesman_problem
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see tsp_test.go

package backtrack

import "math"

// TSP is a solution of the traveling salesman problem
// we assume that the graph is connected
func TSP(graph [][]int) (int, []int) {
	n := len(graph)
	// x is the current solution, bestX is the best solution found so far
	x, bestX := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = i
	} // initialize x as the permutation of the numbers from 0 to n-1
	bestCost := math.MaxInt32 // best cost
	backtrackTsp(graph, x, 1, 0, &bestCost, bestX)
	return bestCost, bestX
}

func backtrackTsp(graph [][]int, x []int, d int, cost int, bestCost *int, bestX []int) {
	if d == len(graph) {
		if cost < *bestCost {
			*bestCost = cost + graph[x[len(graph)-1]][x[0]]
			copy(bestX, x)
		}
		return
	} else if d < len(graph) {
		for i := d; i < len(graph); i++ {
			x[d], x[i] = x[i], x[d]
			// if the cost of the current solution is better
			// than the best solution found so far
			if cost+graph[x[d-1]][x[d]] < *bestCost {
				backtrackTsp(graph, x, d+1, cost+graph[x[d-1]][x[d]], bestCost, bestX)
			}
			x[d], x[i] = x[i], x[d]
		}
	}
}
