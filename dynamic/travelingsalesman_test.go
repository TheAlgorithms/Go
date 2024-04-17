package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestTravelingSalesman(t *testing.T) {
	costTables := [][][]int{
		{
			{0, 10, 15, 20},
			{10, 0, 35, 25},
			{15, 35, 0, 30},
			{20, 25, 30, 0},
		},
		{
			{0, 29, 20, 21},
			{29, 0, 15, 17},
			{20, 15, 0, 28},
			{21, 17, 28, 0},
		},
		{
			{0, 12, 10},
			{12, 0, 15},
			{10, 15, 0},
		},
	}

	expectedResults := []int{
		80,
		73,
		37,
	}

	for i, costTable := range costTables {
		expected := expectedResults[i]
		t.Run(fmt.Sprintf("TSP Case %d", i+1), func(t *testing.T) {
			result := dynamic.TravelingSalesman(costTable)
			if result != expected {
				t.Errorf("Expected %d, but got %d", expected, result)
			}
		})
	}
}
