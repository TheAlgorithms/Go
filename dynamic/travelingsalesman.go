// filename: travelingsalesman.go
// description: The TravelingSalesman function calculates the minimum cost to complete a round-trip through all cities, starting and ending at the first city, using dynamic programming with bitmasking to track visited cities.
// details:
// The code defines a Go function tsp that solves the Traveling Salesman Problem using dynamic programming and bitmasking to track the subsets of visited cities.
// It initializes a DP table where dp[mask][i] holds the minimum cost to visit cities represented by mask and end at city i.
// The function iterates through all possible city subsets to update the DP table.
// Finally determining the least costly route to visit all cities and return to the starting point.
// author [Harshith Sai (SAI)](https://github.com/harshithsaiv)

package dynamic

import (
	"math"
)

const MAX = math.MaxInt32 / 2

func TravelingSalesman(cost [][]int) int {
	n := len(cost)
	allVisited := (1 << n) - 1
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = MAX
		}
	}
	dp[1][0] = 0

	for mask := 1; mask < (1 << n); mask += 2 {
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) != 0 {
				for j := 0; j < n; j++ {
					if (mask & (1 << j)) == 0 {
						newMask := mask | (1 << j)
						if dp[mask][i]+cost[i][j] < dp[newMask][j] {
							dp[newMask][j] = dp[mask][i] + cost[i][j]
						}
					}
				}
			}
		}
	}

	res := MAX
	for i := 1; i < n; i++ {
		if dp[allVisited][i]+cost[i][0] < res {
			res = dp[allVisited][i] + cost[i][0]
		}
	}

	return res
}
