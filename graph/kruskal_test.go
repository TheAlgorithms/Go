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
					Start:  0,
					End:    1,
					Weight: 4,
				},
				{
					Start:  0,
					End:    2,
					Weight: 13,
				},
				{
					Start:  0,
					End:    3,
					Weight: 7,
				},
				{
					Start:  0,
					End:    4,
					Weight: 7,
				},
				{
					Start:  1,
					End:    2,
					Weight: 9,
				},
				{
					Start:  1,
					End:    3,
					Weight: 3,
				},
				{
					Start:  1,
					End:    4,
					Weight: 7,
				},
				{
					Start:  2,
					End:    3,
					Weight: 10,
				},
				{
					Start:  2,
					End:    4,
					Weight: 14,
				},
				{
					Start:  3,
					End:    4,
					Weight: 4,
				},
			},
			cost: 20,
		},
		{
			n: 3,
			graph: []Edge{
				{
					Start:  0,
					End:    1,
					Weight: 12,
				},
				{
					Start:  0,
					End:    2,
					Weight: 18,
				},
				{
					Start:  1,
					End:    2,
					Weight: 6,
				},
			},
			cost: 18,
		},
		{
			n: 4,
			graph: []Edge{
				{
					Start:  0,
					End:    1,
					Weight: 2,
				},
				{
					Start:  0,
					End:    2,
					Weight: 1,
				},
				{
					Start:  0,
					End:    3,
					Weight: 2,
				},
				{
					Start:  1,
					End:    2,
					Weight: 1,
				},
				{
					Start:  1,
					End:    3,
					Weight: 2,
				},
				{
					Start:  2,
					End:    3,
					Weight: 3,
				},
			},
			cost: 4,
		},
		{
			n: 2,
			graph: []Edge{
				{
					Start:  0,
					End:    1,
					Weight: 4000000,
				},
			},
			cost: 4000000,
		},
		{
			n: 1,
			graph: []Edge{
				{
					Start:  0,
					End:    0,
					Weight: 0,
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
