package levenshtein

import "testing"

var testCases = []struct {
	name             string
	string1          string
	string2          string
	insertionCost    int
	substitutionCost int
	deletionCost     int
	expected         int
}{
	{
		"strings with equal operation weights.",
		"stingy",
		"ring",
		1,
		1,
		1,
		3,
	},
	{
		"strings with different operation weights.",
		"stingy",
		"ring",
		1,
		1,
		3,
		7,
	},
	{
		"strings with different operation weights.",
		"kitten",
		"sitting",
		1,
		1,
		1,
		3,
	},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Distance(tc.string1, tc.string2, tc.insertionCost, tc.substitutionCost, tc.deletionCost)
			if actual != tc.expected {
				t.Errorf("Expected Levenshtein distance between strings: '%s' and '%s' is %v, but got: %v", tc.string1, tc.string2, tc.expected, actual)
			}
		})
	}
}
