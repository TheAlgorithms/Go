package graph

import (
	"testing"
)

func TestKahn(t *testing.T) {
	testCases := []struct {
		name         string
		n            int
		dependencies [][]int
		wantNil      bool
	}{
		{
			"linear graph",
			3,
			[][]int{{0, 1}, {1, 2}},
			false,
		},
		{
			"diamond graph",
			4,
			[][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			false,
		},
		{
			"star graph",
			5,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			false,
		},
		{
			"disconnected graph",
			5,
			[][]int{{0, 1}, {0, 2}, {3, 4}},
			false,
		},
		{
			"cycle graph 1",
			4,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			true,
		},
		{
			"cycle graph 2",
			4,
			[][]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}},
			true,
		},
		{
			"single node graph",
			1,
			[][]int{},
			false,
		},
		{
			"empty graph",
			0,
			[][]int{},
			false,
		},
		{
			"redundant dependencies",
			4,
			[][]int{{0, 1}, {1, 2}, {1, 2}, {2, 3}},
			false,
		},
		{
			"island vertex",
			4,
			[][]int{{0, 1}, {0, 2}},
			false,
		},
		{
			"more complicated graph",
			14,
			[][]int{{1, 9}, {2, 0}, {3, 2}, {4, 5}, {4, 6}, {4, 7}, {6, 7},
				{7, 8}, {9, 4}, {10, 0}, {10, 1}, {10, 12}, {11, 13},
				{12, 0}, {12, 11}, {13, 5}},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Kahn(tc.n, tc.dependencies)
			if tc.wantNil {
				if actual != nil {
					t.Errorf("Kahn(%d, %v) = %v; want nil", tc.n, tc.dependencies, actual)
				}
			} else {
				if actual == nil {
					t.Errorf("Kahn(%d, %v) = nil; want valid order", tc.n, tc.dependencies)
				} else {
					seen := make([]bool, tc.n)
					positions := make([]int, tc.n)
					for i, v := range actual {
						seen[v] = true
						positions[v] = i
					}
					for i, v := range seen {
						if !v {
							t.Errorf("missing vertex %v", i)
						}
					}
					for _, d := range tc.dependencies {
						if positions[d[0]] > positions[d[1]] {
							t.Errorf("dependency %v not satisfied", d)
						}
					}
				}
			}
		})
	}
}
