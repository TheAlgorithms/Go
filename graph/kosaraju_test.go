package graph

import (
	"fmt"
	"sort"
	"testing"
)

func TestKosaraju(t *testing.T) {

	var testCases = []struct {
		n   int
		adj [][]int
		SCC [][]int
	}{
		{
			n: 10,
			adj: [][]int{
				{},
				{},
				{},
				{2, 9},
				{5, 10},
				{1},
				{1, 5, 8},
				{8, 6, 10},
				{4, 1, 9},
				{3, 10},
				{7, 8, 5, 9},
			},
			SCC: [][]int{
				{1},
				{2},
				{3, 4, 6, 7, 8, 9, 10},
				{5},
			},
		},
		{
			n: 1,
			adj: [][]int{
				{},
				{},
			},
			SCC: [][]int{
				{1},
			},
		},
	}

	for i := range testCases {

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			components := GetStronglyConnectedComponents(testCases[i].n, testCases[i].adj)

			sort.Slice(components[:], func(j, k int) bool {
				for x := range components[j] {
					if components[j][x] == components[k][x] {
						continue
					}
					return components[j][x] < components[k][x]
				}
				return false
			})

			var failed = false
			if len(components) != len(testCases[i].SCC) {
				failed = true
			}
			if !failed {
				for k := 0; k < len(components); k++ {
					sort.Ints(components[k])
					if len(components[k]) != len(testCases[i].SCC[k]) {
						failed = true
						break
					}
					for l := 0; l < len(components[k]); l++ {
						if components[k][l] != testCases[i].SCC[k][l] {
							failed = true
							break
						}
					}
				}
			}
			if failed {
				t.Error("Components don't match. Expected", testCases[i].SCC, " Found: ", components)
			}
		})
	}
}
