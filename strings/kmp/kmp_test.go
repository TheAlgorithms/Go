package kmp

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	word     string
	text     string
	expected Result
}{
	{
		"String comparison on single pattern match",
		"announce",
		"CPM_annual_conference_announce",
		Result{
			22,
			32,
		},
	},
	{
		"String comparison on multiple pattern match",
		"AABA",
		"AABAACAADAABAABA",
		Result{
			0,
			4,
		},
	},
	{
		"String comparison with not found pattern",
		"AABC",
		"AABAACAADAABAABA",
		Result{
			-1,
			23,
		},
	},
}

func TestKMP(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Kmp(tc.text, tc.word)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected matches for pattern '%s' for string '%s' are: %v steps at position %v, but actual matches are: %v steps at position %v",
					tc.word, tc.text, tc.expected.numberOfComparison, tc.expected.resultPosition, actual.numberOfComparison, actual.resultPosition)
			}
		})
	}
}
