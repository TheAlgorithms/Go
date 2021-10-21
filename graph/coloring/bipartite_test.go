package coloring

import (
	"testing"
)

var testCases = []struct {
	name        string
	N           int
	isBipartite bool
	edges       [][]int
}{
	{
		"basic true", 2, true,
		[][]int{{1, 0}},
	},
	{
		"basic false", 3, false,
		[][]int{{0, 1}, {1, 2}, {2, 0}},
	},
}

func TestBipartite(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BipartiteCheck(tc.N, tc.edges)
			if tc.isBipartite != actual {
				t.Errorf("failed %s: %v", tc.name, tc.edges)
			}
		})
	}
}
