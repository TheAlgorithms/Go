// napsack Problem
// https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/
package napsack

// package main

import (
	"fmt"
	"math"
)

type item struct {
	weight int
	value  int
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// display result of dp data structure
func dist(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		str := ""
		for j := 0; j < len(dp[i]); j++ {
			str = str + " " + fmt.Sprintf("%d", dp[i][j])
		}
		fmt.Println(str)
	}
}

// solve napsack problem
func solve(maxWeight int, items []item) [][]int {
	n := len(items) + 1
	m := maxWeight + 1
	// create dp data structure
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 0; i < len(items); i++ {
		for j := 0; j <= maxWeight; j++ {
			if items[i].weight > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = max(dp[i][j-items[i].weight]+items[i].value, dp[i][j])
			}
		}
	}
	return dp
}

// convert int slice into item slice
func baseConv(base [][]int) []item {
	var items []item
	for i := 0; i < len(base); i++ {
		var tmp item
		tmp.weight = base[i][0]
		tmp.value = base[i][1]
		items = append(items, tmp)
	}
	return items
}

/*
func main() {
	maxWeight := 50
	base := [][]int{
		{10, 60},
		{20, 100},
		{30, 120},
	}
	items := baseConv(base)
	dp := solve(maxWeight, items)
	dist(dp)
}
*/
