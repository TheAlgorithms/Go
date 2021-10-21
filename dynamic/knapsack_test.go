package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestKnapsack(t *testing.T) {
	td := []struct {
		maxWeight int
		weights   []int
		values    []int
		expected  int
	}{
		{0, []int{0}, []int{0}, 0},
		{10, []int{1, 2, 3}, []int{1, 1, 1}, 3},                              // picks all
		{10, []int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 1, 1, 1}, 4},            // picks 1,2,3,4
		{10, []int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 1, 1, 5}, 7},            // picks 1,3,6
		{10, []int{1, 2, 3, 4, 5, 6}, []int{-1, 10, -3, -4, 10, 1}, 20},      // picks 2,5
		{10, []int{1, 2, 3, 4, 5, 6}, []int{-10, -10, -10, -10, 10, 10}, 10}, // picks 5 or 6
	}
	for _, tc := range td {
		name := fmt.Sprintf("Knapsack problem with (maxWeight: %d, weights: %v, values: %v)", tc.maxWeight, tc.weights, tc.values)
		t.Run(name, func(t *testing.T) {
			actual := dynamic.Knapsack(tc.maxWeight, tc.weights, tc.values)
			if actual != tc.expected {
				t.Errorf("expecting knapsack with (maxWeight: %d, weights: %v, values: %v) to return %d but got %d", tc.maxWeight, tc.weights, tc.values, tc.expected, actual)
			}
		})
	}
}

func ExampleKnapsack() {
	fmt.Print(dynamic.Knapsack(10, []int{4, 5, 8}, []int{50, 15, 60}))
	//Output:65
}
