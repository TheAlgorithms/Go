package dynamic

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	testCases := map[string]struct {
		m    int
		n    int
		want int
	}{
		"negative sizes":               {-1, -1, 0},
		"empty matrix both dimensions": {0, 0, 0},
		"empty matrix one dimension":   {0, 1, 0},
		"one element":                  {1, 1, 1},
		"small matrix":                 {2, 2, 2},
		"stress test":                  {1000, 1000, 2874513998398909184},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := UniquePaths(test.m, test.n); got != test.want {
				t.Errorf("UniquePaths(%v, %v) = %v, want %v", test.m, test.n, got, test.want)
			}
		})
	}
}
