// This file provides tests for coloring using Greedy approach.
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring_test

import (
	"strconv"
	"testing"
)

func TestGraphColorUsingGreedyApproach(t *testing.T) {
	for i, tt := range getTestGraphs() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			colorsOfVertices, colors := tt.Graph.ColorUsingGreedyApproach()
			if colors != tt.ColorsUsed {
				t.Errorf("ColorUsingGreedyApproach() return more number of colors: %v, want %v colors", colors, tt.ColorsUsed)
			}
			// check colors
			if err := tt.Graph.ValidateColorsOfVertex(colorsOfVertices); err != nil {
				t.Errorf("ColorUsingGreedyApproach() assigned colors are wrong, error = %v", err)
			}
		})
	}
}
