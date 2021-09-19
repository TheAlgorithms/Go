package coloring_test

import (
	"strconv"
	"testing"
)

func TestGraphColoringUsingBFS(t *testing.T) {
	for i, tt := range getTestGraphs() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			colorsOfVertices, colors := tt.Graph.ColoringUsingBFS()
			if colors != tt.ColorsUsed {
				t.Errorf("ColoringUsingBFS() return more number of colors: %v, want %v colors", colors, tt.ColorsUsed)
			}
			// check colors
			if err := tt.Graph.ValidateColorsOfVertex(colorsOfVertices); err != nil {
				t.Errorf("ColoringUsingBFS() assigned colors are wrong, error = %v", err)
			}
		})
	}
}
