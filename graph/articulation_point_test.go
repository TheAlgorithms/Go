package graph

import (
	"reflect"
	"testing"
)

func TestArticulationPoints(t *testing.T) {
	var testCases = []struct {
		description string
		graph       Graph
		expected    []bool
	}{
		{
			"Linear tree structure",
			Graph{
				vertices: 5,
				edges: map[int]map[int]int{
					0: map[int]int{
						1: 0,
					},

					1: map[int]int{
						2: 0,
					},

					2: map[int]int{
						3: 0,
					},

					3: map[int]int{
						4: 0,
					},
				},
			},
			[]bool{false, true, true, true, false},
		}, {
			"A complete graph",
			Graph{
				vertices: 4,
				edges: map[int]map[int]int{
					0: map[int]int{
						1: 0,
						2: 0,
						3: 0,
					},

					1: map[int]int{
						2: 0,
						3: 0,
					},

					2: map[int]int{
						3: 0,
					},
				},
			},
			[]bool{false, false, false, false},
		},
	}

	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			is_ap := ArticulationPoint(test.graph)
			if reflect.DeepEqual(is_ap, test.expected) {
				t.Logf("FAIL: %s", test.description)
			}
		})
	}
}
