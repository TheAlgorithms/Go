package graph

import (
	"math"
	"testing"
)

//noConnection is a value (INF) which represents that there are no two edges that directly connects these two nodes
var noConnection = math.Inf(1)

func compareDist(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

func equalMatrices(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] == noConnection && b[i][j] == noConnection {
				continue
			}

			if !compareDist(a[i][j], b[i][j]) {
				return false
			}
		}
	}

	return true
}

func TestFloydWarshall(t *testing.T) {
	var floydWarshallTestData = []struct {
		description string
		graph       [][]float64
		expected    [][]float64
	}{
		{
			description: "test empty graph",
			graph:       nil,
			expected:    nil,
		},
		{
			description: "test graph with wrong dimensions",
			graph: [][]float64{
				{1, 2},
				{noConnection},
			},
			expected: nil,
		},
		{
			description: "test graph with no edges",
			graph: [][]float64{
				{noConnection, noConnection},
				{noConnection, noConnection},
			},
			expected: [][]float64{
				{noConnection, noConnection},
				{noConnection, noConnection},
			},
		},
		{
			description: "test graph with only negative edges",
			graph: [][]float64{
				{-3, -2},
				{-3, -2},
			},
			expected: [][]float64{
				{-17, -25},
				{-26, -34},
			},
		},
		{
			description: "test graph with 5 vertices and self-loops",
			graph: [][]float64{
				{1, 2, noConnection, noConnection, noConnection},
				{noConnection, noConnection, 3, -4, noConnection},
				{noConnection, noConnection, noConnection, noConnection, 5},
				{1, noConnection, noConnection, noConnection, noConnection},
				{noConnection, noConnection, noConnection, 2, noConnection},
			},
			expected: [][]float64{
				{-1, 1, 4, -3, 8},
				{-3, -1, 2, -5, 6},
				{7, 9, 12, 5, 5},
				{0, 2, 5, -2, 9},
				{2, 4, 7, 0, 9},
			},
		},
	}
	for _, test := range floydWarshallTestData {
		t.Run(test.description, func(t *testing.T) {
			result := FloydWarshall(test.graph)

			if !equalMatrices(result, test.expected) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expected result:%f\nFound: %f", test.expected, result)
			}
		})
	}
}
