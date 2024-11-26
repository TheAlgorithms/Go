package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestKosaraju(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    map[int][]int
		expected [][]int
	}{
		{
			name:     "Single SCC",
			vertices: 5,
			edges: map[int][]int{
				0: {1},
				1: {2},
				2: {0, 3},
				3: {4},
				4: {},
			},
			expected: [][]int{{4}, {3}, {0, 2, 1}},
		},
		{
			name:     "Multiple SCCs",
			vertices: 8,
			edges: map[int][]int{
				0: {1},
				1: {2},
				2: {0, 3},
				3: {4},
				4: {5},
				5: {3, 6},
				6: {7},
				7: {6},
			},
			expected: [][]int{{6, 7}, {3, 4, 5}, {0, 2, 1}},
		},
		{
			name:     "Disconnected graph",
			vertices: 4,
			edges: map[int][]int{
				0: {1},
				1: {},
				2: {3},
				3: {},
			},
			expected: [][]int{{1}, {0}, {3}, {2}},
		},
		{
			name:     "No edges",
			vertices: 3,
			edges: map[int][]int{
				0: {},
				1: {},
				2: {},
			},
			expected: [][]int{{0}, {1}, {2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initializing graph
			graph := &Graph{
				vertices: tt.vertices,
				edges:    make(map[int]map[int]int),
			}
			for v, neighbors := range tt.edges {
				graph.edges[v] = make(map[int]int)
				for _, neighbor := range neighbors {
					graph.edges[v][neighbor] = 1
				}
			}

			// Running Kosaraju's algorithm to get the SCCs
			result := graph.Kosaraju()

			// Sort the expected and result SCCs to ensure order doesn't matter
			sortSlices(tt.expected)
			sortSlices(result)

			// Compare the sorted SCCs
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Utility function to sort the slices and their contents
func sortSlices(s [][]int) {
	for _, inner := range s {
		sort.Ints(inner)
	}
	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) == 0 || len(s[j]) == 0 {
			return len(s[i]) < len(s[j])
		}
		return s[i][0] < s[j][0]
	})
}
