package graph

import (
	"testing"
)

func TestEdmondKarp(t *testing.T) {
	var edmondKarpTestData = []struct {
		description string
		graph       WeightedGraph
		source      int
		sink        int
		expected    float64
	}{
		{
			description: "test empty graph",
			graph:       nil,
			source:      0,
			sink:        0,
			expected:    0.0,
		},
		{
			description: "test graph with wrong dimensions",
			graph: WeightedGraph{
				{1, 2},
				{0},
			},
			source:   0,
			sink:     1,
			expected: 0.0,
		},
		{
			description: "test graph with no edges",
			graph: WeightedGraph{
				{0, 0},
				{0, 0},
			},
			source:   0,
			sink:     1,
			expected: 0.0,
		},
		{
			description: "test graph with 4 vertices",
			graph: WeightedGraph{
				{0, 1000000, 1000000, 0},
				{0, 0, 1, 1000000},
				{0, 0, 0, 1000000},
				{0, 0, 0, 0},
			},
			source:   0,
			sink:     3,
			expected: 2000000,
		},
		{
			description: "test graph with 6 vertices and some float64 weights",
			graph: WeightedGraph{
				{0, 16, 13.8, 0, 0, 0},
				{0, 0, 10, 12.7, 0, 0},
				{0, 4.2, 0, 0, 14, 0},
				{0, 0, 9.1, 0, 0, 21.3},
				{0, 0, 0, 7.5, 0, 4},
				{0, 0, 0, 0, 0, 0},
			},
			source:   0,
			sink:     5,
			expected: 24.2,
		},
	}
	for _, test := range edmondKarpTestData {
		t.Run(test.description, func(t *testing.T) {
			result := EdmondKarp(test.graph, test.source, test.sink)

			if !almostEqual(test.expected, result) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expected result:%f\nFound: %f", test.expected, result)
			}
		})
	}
}
