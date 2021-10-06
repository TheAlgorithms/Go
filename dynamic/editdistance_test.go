package dynamic

import (
	"fmt"
	"testing"
)

func Test_EditDistance(t *testing.T) {

	var testCases = []struct {
		first    string
		second   string
		expected int
	}{
		{"", "", 0},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"abcdxabcde", "abcdeabcdx", 2},
		{"sunday", "saturday", 3},
		{"food", "money", 4},
		{"voldemort", "dumbledore", 7},
	}

	for i := range testCases {

		t.Run(fmt.Sprintf("Word 1: %s, Word 2: %s", testCases[i].first, testCases[i].second), func(t *testing.T) {

			computed := EditDistanceDP(testCases[i].first, testCases[i].second)

			if computed != testCases[i].expected {
				t.Errorf("Word 1: %s, Word 2: %s, Expected: %d, Computed: %d", testCases[i].first, testCases[i].second, testCases[i].expected, computed)
			}
		})
	}
}
