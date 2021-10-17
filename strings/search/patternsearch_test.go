package search

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	input    string
	pattern  string
	expected []int
}{
	{
		"string with multiple pattern matches",
		"ABAAABCDBBABCDDEBCABC",
		"ABC",
		[]int{4, 10, 18},
	},
	{
		"string with single pattern match",
		"ABCDEFGHIJKLMN",
		"CDE",
		[]int{2},
	},
	{
		"string with no pattern match",
		"ABCDEFGHIJKLMN",
		"XYZ",
		[]int(nil),
	},
	{
		"empty string",
		"",
		"XYZ",
		[]int(nil),
	},
	{
		"worse case 1",
		"AAAAAAAAAA",
		"AAA",
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		"worse case 2",
		"NANANANANANANANANA",
		"NANANA",
		[]int{0, 2, 4, 6, 8, 10, 12},
	},
}

func TestNaive(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Naive(tc.input, tc.pattern)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected matches for pattern '%s' for string '%s' are: %v, but actual matches are: %v", tc.pattern, tc.input, tc.expected, actual)
			}
		})
	}
}

func TestBooyerMoore(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BoyerMoore(tc.input, tc.pattern)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected matches for pattern '%s' for string '%s' are: %v, but actual matches are: %v", tc.pattern, tc.input, tc.expected, actual)
			}
		})
	}
}
