package graph

import (
	"testing"
)

var testCases = []struct {
	name        string
	N           int
	constraints [][]int
}{
	{
		"basic test", 2,
		[][]int{{1, 0}},
	},
	{
		"double path", 7,
		[][]int{
			{0, 1}, {1, 3}, {3, 5},
			{0, 2}, {2, 4}, {4, 6}},
	},
	{
		"star shape", 7,
		[][]int{
			{0, 1}, {0, 3}, {0, 5},
			{0, 2}, {0, 4}, {0, 6}},
	},
	{
		"tree shape", 7,
		[][]int{
			{0, 1}, {1, 3}, {1, 5},
			{0, 2}, {2, 4}, {2, 6}},
	},
}

func TestTopological(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Topological(tc.N, tc.constraints)

			visited := make([]bool, tc.N)
			positions := make([]int, tc.N)
			for i := 0; i < tc.N; i++ {
				positions[actual[i]] = i
				visited[actual[i]] = true
			}
			for _, v := range visited {
				if !v {
					t.Errorf("nodes not all visited, %v", visited)
				}
			}
			for _, c := range tc.constraints {
				if positions[c[0]] > positions[c[1]] {
					t.Errorf("%v dun satisfy %v", actual, c)
				}
			}
		})
	}
}
