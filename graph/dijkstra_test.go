// dijkstra_test.go
// description: Comprehensive tests for the improved Dijkstra algorithm implementation
// details:
// This file contains extensive test cases for all variants of the Dijkstra algorithm:
// - Basic shortest path finding
// - Path reconstruction functionality
// - All-pairs shortest paths
// - Edge cases and error handling
// - Performance benchmarks
//
// Test coverage includes:
// - Connected and disconnected graphs
// - Self-loops and single nodes
// - Invalid inputs and error conditions
// - Large graphs for performance testing
// - Various graph topologies (linear, star, complete, etc.)
//
// author(s) [Claude](https://github.com/anthropic)
// see dijkstra.go, graph.go

package graph

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// dijkstraTestCase represents a test case for Dijkstra algorithm testing
type dijkstraTestCase struct {
	name          string  // Test case description
	edges         [][]int // Graph edges in format [from, to, weight]
	start         int     // Starting node
	end           int     // Ending node
	expectedDist  int     // Expected shortest distance
	expectedFound bool    // Expected result for path existence
	expectedPath  []int   // Expected path (for path reconstruction tests)
}

// Test cases for basic Dijkstra functionality
var dijkstraBasicTests = []dijkstraTestCase{
	{
		name:          "direct connection",
		edges:         [][]int{{0, 1, 5}},
		start:         0,
		end:           1,
		expectedDist:  5,
		expectedFound: true,
		expectedPath:  []int{0, 1},
	},
	{
		name:          "two hop path",
		edges:         [][]int{{0, 1, 5}, {1, 2, 2}},
		start:         0,
		end:           2,
		expectedDist:  7,
		expectedFound: true,
		expectedPath:  []int{0, 1, 2},
	},
	{
		name:          "choose shorter path",
		edges:         [][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}},
		start:         0,
		end:           3,
		expectedDist:  9,
		expectedFound: true,
		expectedPath:  []int{0, 2, 3},
	},
	{
		name:          "longer path with more hops",
		edges:         [][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}, {3, 4, 1}},
		start:         0,
		end:           4,
		expectedDist:  10,
		expectedFound: true,
		expectedPath:  []int{0, 2, 3, 4},
	},
	{
		name:          "unreachable node",
		edges:         [][]int{{0, 1, 5}},
		start:         0,
		end:           2,
		expectedDist:  -1,
		expectedFound: false,
		expectedPath:  nil,
	},
	{
		name:          "self loop",
		edges:         [][]int{{0, 0, 1}},
		start:         0,
		end:           0,
		expectedDist:  0,
		expectedFound: true,
		expectedPath:  []int{0},
	},
	{
		name: "complex graph with multiple paths",
		edges: [][]int{
			{0, 1, 4}, {0, 2, 1}, {1, 3, 1}, {2, 1, 2}, {2, 3, 5}, {3, 4, 3},
		},
		start:         0,
		end:           4,
		expectedDist:  7, // Corrected: 0->2(1) + 2->1(2) + 1->3(1) + 3->4(3) = 7
		expectedFound: true,
		expectedPath:  []int{0, 2, 1, 3, 4},
	},
	{
		name:          "triangle with optimal path",
		edges:         [][]int{{0, 1, 10}, {0, 2, 3}, {2, 1, 4}},
		start:         0,
		end:           1,
		expectedDist:  7,
		expectedFound: true,
		expectedPath:  []int{0, 2, 1},
	},
}

// Edge case tests for error handling and special scenarios
var dijkstraEdgeCaseTests = []dijkstraTestCase{
	{
		name:          "empty graph",
		edges:         [][]int{},
		start:         0,
		end:           1,
		expectedDist:  -1,
		expectedFound: false,
		expectedPath:  nil,
	},
	{
		name:          "single node graph",
		edges:         [][]int{},
		start:         0,
		end:           0,
		expectedDist:  -1, // Node doesn't exist
		expectedFound: false,
		expectedPath:  nil,
	},
	{
		name:          "negative start node",
		edges:         [][]int{{0, 1, 5}},
		start:         -1,
		end:           1,
		expectedDist:  -1,
		expectedFound: false,
		expectedPath:  nil,
	},
	{
		name:          "large weights",
		edges:         [][]int{{0, 1, 1000000}, {1, 2, 1000000}},
		start:         0,
		end:           2,
		expectedDist:  2000000,
		expectedFound: true,
		expectedPath:  []int{0, 1, 2},
	},
}

// createTestGraph creates a graph from edge specifications for testing
func createTestGraph(edges [][]int) *Graph {
	// Find max vertex to determine initial size
	maxVertex := -1
	for _, edge := range edges {
		if edge[0] > maxVertex {
			maxVertex = edge[0]
		}
		if edge[1] > maxVertex {
			maxVertex = edge[1]
		}
	}

	// Create graph using existing constructor
	g := New(maxVertex + 1)

	// Add edges using existing method
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		g.AddWeightedEdge(from, to, weight)
	}

	return g
}

// TestDijkstraBasic tests basic Dijkstra functionality
func TestDijkstraBasic(t *testing.T) {
	for _, test := range dijkstraBasicTests {
		t.Run(test.name, func(t *testing.T) {
			graph := createTestGraph(test.edges)

			distance, found := graph.Dijkstra(test.start, test.end)

			if found != test.expectedFound {
				t.Errorf("Expected found=%v, got found=%v", test.expectedFound, found)
			}

			if distance != test.expectedDist {
				t.Errorf("Expected distance=%d, got distance=%d", test.expectedDist, distance)
			}
		})
	}
}

// TestDijkstraEdgeCases tests edge cases and error conditions
func TestDijkstraEdgeCases(t *testing.T) {
	for _, test := range dijkstraEdgeCaseTests {
		t.Run(test.name, func(t *testing.T) {
			graph := createTestGraph(test.edges)

			distance, found := graph.Dijkstra(test.start, test.end)

			if found != test.expectedFound {
				t.Errorf("Expected found=%v, got found=%v", test.expectedFound, found)
			}

			if distance != test.expectedDist {
				t.Errorf("Expected distance=%d, got distance=%d", test.expectedDist, distance)
			}
		})
	}
}

// TestDijkstraNilGraph tests behavior with nil graph
func TestDijkstraNilGraph(t *testing.T) {
	var g *Graph = nil

	distance, found := g.Dijkstra(0, 1)

	if found != false {
		t.Errorf("Expected found=false for nil graph, got found=%v", found)
	}

	if distance != -1 {
		t.Errorf("Expected distance=-1 for nil graph, got distance=%d", distance)
	}
}

// TestDijkstraWithPath tests path reconstruction functionality
func TestDijkstraWithPath(t *testing.T) {
	for _, test := range dijkstraBasicTests {
		t.Run(test.name+" with path", func(t *testing.T) {
			graph := createTestGraph(test.edges)

			path, distance, err := graph.DijkstraWithPath(test.start, test.end)

			// Check error conditions
			if !test.expectedFound && err == nil && len(path) > 0 {
				t.Errorf("Expected no path but got path: %v", path)
			}

			if test.expectedFound && len(path) == 0 {
				t.Errorf("Expected path but got empty path")
			}

			// Check distance consistency
			if test.expectedFound && distance != test.expectedDist {
				t.Errorf("Expected distance=%d, got distance=%d", test.expectedDist, distance)
			}

			// Check path validity (if expected path is provided)
			if test.expectedPath != nil && test.expectedFound {
				if !reflect.DeepEqual(path, test.expectedPath) {
					t.Errorf("Expected path=%v, got path=%v", test.expectedPath, path)
				}
			}

			// Verify path correctness by checking edges exist
			if len(path) > 1 {
				totalDist := 0
				for i := 0; i < len(path)-1; i++ {
					from, to := path[i], path[i+1]
					if weight, exists := graph.edges[from][to]; exists {
						totalDist += weight
					} else {
						t.Errorf("Path contains non-existent edge: %d -> %d", from, to)
					}
				}

				if totalDist != distance {
					t.Errorf("Path distance mismatch: expected=%d, calculated=%d", distance, totalDist)
				}
			}
		})
	}
}

// TestDijkstraWithPathErrors tests error handling in path reconstruction
func TestDijkstraWithPathErrors(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph
		start         int
		end           int
		expectError   bool
		errorContains string
	}{
		{
			name:          "nil graph",
			graph:         nil,
			start:         0,
			end:           1,
			expectError:   true,
			errorContains: "nil",
		},
		{
			name:          "nonexistent start node",
			graph:         createTestGraph([][]int{{0, 1, 5}}),
			start:         99,
			end:           1,
			expectError:   true,
			errorContains: "start node does not exist",
		},
		{
			name:          "nonexistent end node",
			graph:         createTestGraph([][]int{{0, 1, 5}}),
			start:         0,
			end:           99,
			expectError:   true,
			errorContains: "end node does not exist",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, err := test.graph.DijkstraWithPath(test.start, test.end)

			if test.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !test.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if test.expectError && err != nil && test.errorContains != "" {
				if !contains(err.Error(), test.errorContains) {
					t.Errorf("Expected error to contain '%s', got: %v", test.errorContains, err)
				}
			}
		})
	}
}

// TestDijkstraAllPairs tests all-pairs shortest path functionality
func TestDijkstraAllPairs(t *testing.T) {
	// Test with a simple connected graph
	t.Run("connected graph", func(t *testing.T) {
		edges := [][]int{{0, 1, 5}, {1, 2, 2}, {0, 2, 8}}
		graph := createTestGraph(edges)

		distances, err := graph.DijkstraAllPairs(0)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		expected := map[int]int{0: 0, 1: 5, 2: 7}

		if !reflect.DeepEqual(distances, expected) {
			t.Errorf("Expected distances=%v, got distances=%v", expected, distances)
		}
	})

	// Test with disconnected graph
	t.Run("disconnected graph", func(t *testing.T) {
		edges := [][]int{{0, 1, 5}, {2, 3, 3}}
		graph := createTestGraph(edges)

		distances, err := graph.DijkstraAllPairs(0)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Should only contain distances to reachable nodes
		expected := map[int]int{0: 0, 1: 5}

		if !reflect.DeepEqual(distances, expected) {
			t.Errorf("Expected distances=%v, got distances=%v", expected, distances)
		}
	})
}

// TestDijkstraAllPairsErrors tests error handling in all-pairs functionality
func TestDijkstraAllPairsErrors(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph
		start         int
		expectError   bool
		errorContains string
	}{
		{
			name:          "nil graph",
			graph:         nil,
			start:         0,
			expectError:   true,
			errorContains: "nil",
		},
		{
			name:          "nonexistent start node",
			graph:         createTestGraph([][]int{{0, 1, 5}}),
			start:         99,
			expectError:   true,
			errorContains: "start node does not exist",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.graph.DijkstraAllPairs(test.start)

			if test.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !test.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if test.expectError && err != nil && test.errorContains != "" {
				if !contains(err.Error(), test.errorContains) {
					t.Errorf("Expected error to contain '%s', got: %v", test.errorContains, err)
				}
			}
		})
	}
}

// TestDijkstraPerformance tests performance with larger graphs
func TestDijkstraPerformance(t *testing.T) {
	// Create a larger test graph (star topology for predictable behavior)
	edges := make([][]int, 0)
	center := 0
	numNodes := 100

	// Connect center to all other nodes
	for i := 1; i < numNodes; i++ {
		edges = append(edges, []int{center, i, i}) // weight = node number for easy verification
	}

	graph := createTestGraph(edges)

	t.Run("large graph shortest path", func(t *testing.T) {
		distance, found := graph.Dijkstra(0, 50)

		if !found {
			t.Errorf("Expected to find path in connected graph")
		}

		if distance != 50 { // Direct connection from center
			t.Errorf("Expected distance=50, got distance=%d", distance)
		}
	})

	t.Run("large graph all pairs", func(t *testing.T) {
		distances, err := graph.DijkstraAllPairs(0)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if len(distances) != numNodes {
			t.Errorf("Expected %d distances, got %d", numNodes, len(distances))
		}

		// Verify some distances
		if distances[0] != 0 {
			t.Errorf("Expected distance to self = 0, got %d", distances[0])
		}

		if distances[25] != 25 {
			t.Errorf("Expected distance to node 25 = 25, got %d", distances[25])
		}
	})
}

// TestDijkstraStressOverflow tests behavior with potential integer overflow
func TestDijkstraStressOverflow(t *testing.T) {
	// Test with very large weights that could cause overflow
	edges := [][]int{
		{0, 1, math.MaxInt32 - 1000},
		{1, 2, 999}, // This should not cause overflow
	}

	graph := createTestGraph(edges)

	distance, found := graph.Dijkstra(0, 2)

	if !found {
		t.Errorf("Expected to find path")
	}

	expected := math.MaxInt32 - 1
	if distance != expected {
		t.Errorf("Expected distance=%d, got distance=%d", expected, distance)
	}
}

// contains checks if a string contains a substring (helper function)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(len(substr) == 0 ||
			(len(s) > 0 && (s[:len(substr)] == substr ||
				(len(s) > len(substr) && contains(s[1:], substr)))))
}

// Benchmark tests for performance measurement
func BenchmarkDijkstra(b *testing.B) {
	// Create test graphs of different sizes
	sizes := []int{10, 50, 100, 500}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("DijkstraSize%d", size), func(b *testing.B) {
			// Create complete graph
			edges := make([][]int, 0)
			for i := 0; i < size; i++ {
				for j := i + 1; j < size; j++ {
					edges = append(edges, []int{i, j, i + j + 1})
				}
			}

			graph := createTestGraph(edges)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				graph.Dijkstra(0, size-1)
			}
		})
	}
}

func BenchmarkDijkstraAllPairs(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("AllPairsSize%d", size), func(b *testing.B) {
			// Create star graph (more realistic for all-pairs)
			edges := make([][]int, 0)
			for i := 1; i < size; i++ {
				edges = append(edges, []int{0, i, i})
			}

			graph := createTestGraph(edges)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				graph.DijkstraAllPairs(0)
			}
		})
	}
}
