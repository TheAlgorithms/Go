package graph

import (
	"fmt"
	"testing"
)

func Test_KruskalMST(t *testing.T) {

	var testCases = []struct {
		n     int
		graph []Edge
		cost  int
	}{
		{
			n: 5,
			graph: []Edge{
				{
					start:  0,
					end:    1,
					weight: 4,
				},
				{
					start:  0,
					end:    2,
					weight: 13,
				},
				{
					start:  0,
					end:    3,
					weight: 7,
				},
				{
					start:  0,
					end:    4,
					weight: 7,
				},
				{
					start:  1,
					end:    2,
					weight: 9,
				},
				{
					start:  1,
					end:    3,
					weight: 3,
				},
				{
					start:  1,
					end:    4,
					weight: 7,
				},
				{
					start:  2,
					end:    3,
					weight: 10,
				},
				{
					start:  2,
					end:    4,
					weight: 14,
				},
				{
					start:  3,
					end:    4,
					weight: 4,
				},
			},
			cost: 20,
		},
		{
			n: 3,
			graph: []Edge{
				{
					start:  0,
					end:    1,
					weight: 12,
				},
				{
					start:  0,
					end:    2,
					weight: 18,
				},
				{
					start:  1,
					end:    2,
					weight: 6,
				},
			},
			cost: 18,
		},
		{
			n: 4,
			graph: []Edge{
				{
					start:  0,
					end:    1,
					weight: 2,
				},
				{
					start:  0,
					end:    2,
					weight: 1,
				},
				{
					start:  0,
					end:    3,
					weight: 2,
				},
				{
					start:  1,
					end:    2,
					weight: 1,
				},
				{
					start:  1,
					end:    3,
					weight: 2,
				},
				{
					start:  2,
					end:    3,
					weight: 3,
				},
			},
			cost: 4,
		},
		{
			n: 2,
			graph: []Edge{
				{
					start:  0,
					end:    1,
					weight: 4000000,
				},
			},
			cost: 4000000,
		},
		{
			n: 1,
			graph: []Edge{
				{
					start:  0,
					end:    0,
					weight: 0,
				},
			},
			cost: 0,
		},
	}

	for i := range testCases {

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {

			_, computed := KruskalMST(testCases[i].n, testCases[i].graph)
			if computed != testCases[i].cost {
				t.Errorf("Test Case %d, Expected: %d, Computed: %d", i, testCases[i].cost, computed)
			}
		})
	}
}
