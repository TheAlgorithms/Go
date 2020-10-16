package bom

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	word     string
	text     string
	expected result
}{
	{
		"String comparison on single pattern match",
		"announce",
		"CPM_annual_conference_announce",
		result{
			1,
			[]int{
				22,
			},
		},
	},
	{
		"String comparison on multiple pattern match",
		"announce",
		"CPM_announceannual_conferenceannounce_announce",
		result{
			3,
			[]int{
				4, 29, 38,
			},
		},
	},
	{
		"String comparison with not found pattern",
		"announcement",
		"CPM_announceannual_conferenceannounce_announce",
		result{
			0,
			[]int{},
		},
	},
}

func TestBOM(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := bom(tc.text, tc.word)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected matches for pattern '%s' for string '%s' are: found %v times at %v, but actual matches are: found %v times at %v",
					tc.word, tc.text, tc.expected.numberOfFoundWord, tc.expected.foundPositions, actual.numberOfFoundWord, actual.foundPositions)
			}
		})
	}
}
