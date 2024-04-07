package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestTrapRainWater(t *testing.T) {
	heights := [][]int{
		{},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{3, 1, 2, 4, 0, 1, 3, 2, 4},
	}

	expectedResults := []int{
		0,
		6,
		9,
		13,
	}

	for i, height := range heights {
		expected := expectedResults[i]
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := dynamic.TrapRainWater(height)
			if result != expected {
				t.Errorf("Expected %d, but got %d", expected, result)
			}
		})
	}
}
