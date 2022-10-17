package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func lpsTestTemplate(t *testing.T, algorithm func(input string) int) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"BBABCBCAB", 7},
		{"BBBAB", 4},
		{"ABBD", 2},
		{"GEEKSFORGEEKS", 5},
		{"abcdefgh", 1},
		{"bbbab", 4},
		{"cbbd", 2},
		{"racexyzcxar", 7},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint("test with ", tc.input), func(t *testing.T) {
			result := algorithm(tc.input)
			if tc.expected != result {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestLpsRec(t *testing.T) {
	lpsTestTemplate(t, dynamic.LpsRec)
}

func TestLpsDp(t *testing.T) {
	lpsTestTemplate(t, dynamic.LpsDp)
}
