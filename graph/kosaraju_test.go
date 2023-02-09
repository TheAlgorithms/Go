package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKosaraju(t *testing.T) {
	tests := []struct {
		name        string
		edges       map[int][]int
		expectedSCC int
	}{
		{
			name: "Compute Kosaraju's algorithm to find number of Strong connected components",
			edges: map[int][]int{0: {1, 5}, 3: {5, 2, 3}, 2: {0, 3},
				5: {4}, 4: {3, 2}, 6: {4, 0, 8, 9}, 8: {6, 7},
				7: {6, 9}, 11: {4, 12}, 9: {11, 10}, 10: {12},
				12: {9}},
			expectedSCC: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			kosarajuImplementation := NewGraphK()
			kosarajuImplementation.PopulateGraph(test.edges)
			scc := kosarajuImplementation.EvaluateSCC()
			assert.Equal(t, len(scc), test.expectedSCC)
		})
	}
}
