package graph

import (
	"reflect"
	"testing"
)

func TestDfsWhenPathIsFound(t *testing.T) {
	nodes := []int{
		1, 2, 3, 4, 5, 6,
	}

	//Adjacency Matrix for connected nodes
	edges := [][]bool{
		{false, true, true, false, false, false},
		{true, false, false, true, false, false},
		{true, false, false, true, false, false},
		{false, true, true, false, true, false},
		{false, false, false, true, false, true},
		{false, false, false, false, true, false},
	}

	start := 1
	end := 6

	actual, actualIsFound := DepthFirstSearch(start, end, nodes, edges)
	expected := []int{1, 3, 4, 5, 6}
	expectedIsFound := true
	t.Run("Test Dfs", func(t *testing.T) {
		if !reflect.DeepEqual(expected, actual) || !reflect.DeepEqual(actualIsFound, expectedIsFound) {
			t.Errorf("got route: %v, want route: %v", actual, expected)
			t.Errorf("got isFound: %v, want isFound: %v", actualIsFound, expectedIsFound)
		}
	})
}

func TestDfsWhenPathIsNotFound(t *testing.T) {
	nodes := []int{
		1, 2, 3, 4, 5, 6,
	}

	//Adjacency Matrix for connected nodes
	edges := [][]bool{
		{false, true, true, false, false, false},
		{true, false, false, true, false, false},
		{true, false, false, true, false, false},
		{false, true, true, false, true, false},
		{false, false, false, true, false, true},
		{false, false, false, false, true, false},
	}

	start := 1
	end := 7

	actual, actualIsFound := DepthFirstSearch(start, end, nodes, edges)
	var expected []int
	expectedIsFound := false
	t.Run("Test Dfs", func(t *testing.T) {
		if !reflect.DeepEqual(expected, actual) || !reflect.DeepEqual(actualIsFound, expectedIsFound) {
			t.Errorf("got route: %v, want route: %v", actual, expected)
			t.Errorf("got isFound: %v, want isFound: %v", actualIsFound, expectedIsFound)
		}
	})
}
