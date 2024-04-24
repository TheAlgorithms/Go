// filename: travelingsalesman.go
// description: The TravelingSalesman function calculates the minimum cost to complete a round-trip through all cities, starting and ending at the first city, using dynamic programming with bitmasking to track visited cities.
// details:
// The code defines a Go function tsp that solves the Traveling Salesman Problem using dynamic programming and bitmasking to track the subsets of visited cities.
// It initializes a DP table where dp[mask][i] holds the minimum cost to visit cities represented by mask and end at city i.
// The function iterates through all possible city subsets to update the DP table.
// Finally determining the least costly route to visit all cities and return to the starting point.
// author [Harshith Sai (SAI)](https://github.com/harshithsaiv)

package dynamic

import "math"

const maxCost = math.MaxInt32 / 2

// initializeCostTable creates and initializes the DP table with maximum costs.
func initializeCostTable(numCities int) [][]int {
	costTable := make([][]int, 1<<numCities)
	for i := range costTable {
		costTable[i] = make([]int, numCities)
		for j := range costTable[i] {
			costTable[i][j] = maxCost
		}
	}
	costTable[1][0] = 0
	return costTable
}

// updateCostTable updates the DP table with the minimum cost from `current` to `next` city.
func updateCostTable(costTable [][]int, mask, numCities int, cost [][]int) {
	for current := 0; current < numCities; current++ {
		if (mask & (1 << current)) == 0 {
			continue
		}
		for next := 0; next < numCities; next++ {
			if (mask&(1<<next)) != 0 || current == next {
				continue
			}
			newMask := mask | (1 << next)
			newCost := costTable[mask][current] + cost[current][next]
			if newCost < costTable[newMask][next] {
				costTable[newMask][next] = newCost
			}
		}
	}
}

// findMinimumCost calculates the minimal return cost to the starting city.
func findMinimumCost(costTable [][]int, allVisited, numCities int, cost [][]int) int {
	minCost := maxCost
	for i := 1; i < numCities; i++ {
		finalCost := costTable[allVisited][i] + cost[i][0]
		if finalCost < minCost {
			minCost = finalCost
		}
	}
	return minCost
}

// TravelingSalesman calculates the minimum cost to complete a round trip through all cities.
func TravelingSalesman(cost [][]int) int {
	numCities := len(cost)
	allVisited := (1 << numCities) - 1
	costTable := initializeCostTable(numCities)
	for mask := 1; mask < (1 << numCities); mask += 2 {
		updateCostTable(costTable, mask, numCities, cost)
	}
	return findMinimumCost(costTable, allVisited, numCities, cost)
}
