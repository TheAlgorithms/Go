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
	var tsp func(int, int)
	tsp = func(d int, cost int) { // depth, cost so far
		if d == n {
			if cost < bestCost {
				bestCost = cost + graph[x[n-1]][x[0]]
				copy(bestX, x)
			}
			return
		} else if d < n {
			for i := d; i < n; i++ {
				x[d], x[i] = x[i], x[d]
				// if the cost of the current solution is better than the best solution found so far
				if cost+graph[x[d-1]][x[d]] < bestCost {
					tsp(d+1, cost+graph[x[d-1]][x[d]])
				}
				x[d], x[i] = x[i], x[d]
			}
		}
	}
	tsp(1, 0) // make start of the tour is the city 0
	return bestCost, bestX
}
