package horspool

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
			22,
			12,
		},
	},
	{
		"String comparison on multiple pattern match",
		"announce",
		"CPM_announceannual_conferenceannounce_announce",
		result{
			4,
			9,
		},
	},
	{
		"String comparison with not found pattern",
		"announcement",
		"CPM_announceannual_conferenceannounce_announce",
		result{
			-1,
			5,
		},
	},
}

func TestHorspool(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := horspool(tc.text, tc.word)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected matches for pattern '%s' for string '%s' are: found at %d with comparison %d times, but actual matches are: found at %d with comparison %d times",
					tc.word, tc.text, tc.expected.foundPosition, tc.expected.numberOfComparison, actual.foundPosition, actual.numberOfComparison)
			}
		})
	}
}
