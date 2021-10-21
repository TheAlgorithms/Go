package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestBin2(t *testing.T) {
	td := []struct {
		n, k, expected int
	}{
		{0, 0, 1},
		{1, 1, 1},
		{2, 0, 1}, {2, 1, 2}, {2, 2, 1},
		{3, 0, 1}, {3, 1, 3}, {3, 2, 3}, {3, 3, 1},
		{4, 0, 1}, {4, 1, 4}, {4, 2, 6}, {4, 3, 4}, {4, 4, 1},
		{5, 0, 1}, {5, 1, 5}, {5, 2, 10}, {5, 3, 10}, {5, 4, 5}, {5, 5, 1},
		{10, 2, 45},
		{15, 10, 3003},
	}
	for _, tc := range td {
		name := fmt.Sprintf("binomial coefficient of (%d, %d)", tc.n, tc.k)
		t.Run(name, func(t *testing.T) {
			actual := dynamic.Bin2(tc.n, tc.k)
			if actual != tc.expected {
				t.Errorf("expecting binomial coefficient of (%d, %d) to be %d but got %d", tc.n, tc.k, tc.expected, actual)
			}
		})
	}
}
