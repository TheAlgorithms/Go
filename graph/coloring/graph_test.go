// This file provides tests for graph coloring validations.
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring_test

import (
	"github.com/TheAlgorithms/Go/graph/coloring"
	"strconv"
	"testing"
)

type testGraph struct {
	Graph        *coloring.Graph
	ColorsUsed   int
	VertexColors map[int]coloring.Color
}

func getTestGraphs() (list []*testGraph) {
	// Graph 0th:
	//     1---2
	//     | /  \
	// 4---3    0
	// Min number of colors required = 3
	g0 := &testGraph{
		Graph:      &coloring.Graph{},
		ColorsUsed: 3,
		VertexColors: map[int]coloring.Color{
			1: 1, 4: 1, 0: 1,
			2: 2,
			3: 3,
		},
	}
	list = append(list, g0)
	g0.Graph.AddEdge(4, 3)
	g0.Graph.AddEdge(3, 1)
	g0.Graph.AddEdge(3, 2)
	g0.Graph.AddEdge(1, 2)
	g0.Graph.AddEdge(2, 0)

	// Graph 1st:
	//      1---2
	//      | / |
	//  4---3---0
	// Min number of colors required = 3
	g1 := &testGraph{
		Graph:      &coloring.Graph{},
		ColorsUsed: 3,
		VertexColors: map[int]coloring.Color{
			1: 1, 4: 1, 0: 1,
			2: 2,
			3: 3,
		},
	}
	list = append(list, g1)
	g1.Graph.AddEdge(4, 3)
	g1.Graph.AddEdge(3, 1)
	g1.Graph.AddEdge(3, 2)
	g1.Graph.AddEdge(1, 2)
	g1.Graph.AddEdge(2, 0)
	g1.Graph.AddEdge(3, 0)

	// Graph 2nd:
	//      1---2
	//      |
	//  4---3   0
	// Min number of colors required = 2
	g2 := &testGraph{
		Graph:      &coloring.Graph{},
		ColorsUsed: 2,
		VertexColors: map[int]coloring.Color{
			1: 1, 4: 1, 0: 1,
			2: 2, 3: 2,
		},
	}
	list = append(list, g2)
	g2.Graph.AddVertex(0)
	g2.Graph.AddEdge(4, 3)
	g2.Graph.AddEdge(3, 1)
	g2.Graph.AddEdge(1, 2)

	// Graph 3rd:
	//  1---2   4
	//  |   |   |
	//  0---3   5
	// Min number of colors required = 2
	g3 := &testGraph{
		Graph:      &coloring.Graph{},
		ColorsUsed: 2,
		VertexColors: map[int]coloring.Color{
			1: 1, 3: 1, 4: 1,
			0: 2, 2: 2, 5: 2,
		},
	}
	list = append(list, g3)
	g3.Graph.AddEdge(0, 3)
	g3.Graph.AddEdge(1, 2)
	g3.Graph.AddEdge(1, 0)
	g3.Graph.AddEdge(3, 2)
	g3.Graph.AddEdge(4, 5)

	// Graph 4th:
	// Completely Connected graph of vertex 4
	// Min number of colors required = 2
	g4 := &testGraph{
		Graph:      &coloring.Graph{},
		ColorsUsed: 4,
		VertexColors: map[int]coloring.Color{
			0: 1, 1: 2, 2: 3, 4: 4,
		},
	}
	list = append(list, g4)
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			g4.Graph.AddEdge(i, j)
		}
	}

	return
}

func TestGraph_ValidateColorsOfVertex(t *testing.T) {
	for i, tt := range getTestGraphs() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if err := tt.Graph.ValidateColorsOfVertex(tt.VertexColors); err != nil {
				t.Errorf("ValidateColorsOfVertex() error = %v, wantErr nil", err)
			}
		})
	}
}

func getTestGraphsForNegativeTests() (list []*testGraph) {
	list = getTestGraphs()
	list[0].VertexColors = nil
	list[1].VertexColors = map[int]coloring.Color{}
	for v := range list[2].VertexColors {
		in := len(list[2].VertexColors) - v - 1
		list[2].VertexColors[v] = list[2].VertexColors[in]
	}
	for v := range list[3].VertexColors {
		list[3].VertexColors[v] = 1
	}
	return list[:4]
}

func TestGraphValidateColorsOfVertex_Negative(t *testing.T) {
	for i, tt := range getTestGraphsForNegativeTests() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if err := tt.Graph.ValidateColorsOfVertex(tt.VertexColors); err == nil {
				t.Errorf("ValidateColorsOfVertex() error = nil, want some err")
			}
		})
	}
}
