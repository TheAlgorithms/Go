package advancedahocorasick

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var testCases = []struct {
	name     string
	words    []string
	text     string
	expected Result
}{

	{
		"String comparison on all patterns found",
		[]string{"announce", "annual", "annually"},
		"CPM_annual_conferenceannounce_announce_annually_announce",
		Result{
			map[string][]int{
				"annual":   {4, 39},
				"announce": {21, 30, 48},
				"annually": {39},
			},
		},
	},
	{
		"String comparison on not all patterns found",
		[]string{"announce", "annual", "annually"},
		"CPM_annual_conference_announce",
		Result{
			map[string][]int{
				"annual":   {4},
				"announce": {22},
			},
		},
	},
	{
		"String comparison on not all patterns found",
		[]string{"announce", "annual", "annually"},
		"CPM_annual_conference_announce",
		Result{
			map[string][]int{
				"annual":   {4},
				"announce": {22},
			},
		},
	},
}

func TestAhoCorasick(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := AhoCorasick(tc.text, tc.words)
			if !reflect.DeepEqual(actual, tc.expected) {
				actualString := convertToString(actual)
				expectedString := convertToString(tc.expected)
				t.Errorf("Expected matches for patterns %s for string '%s' are: patterns and positions found %v, but actual matches are: patterns and positions found %v",
					tc.words, tc.text, actualString, expectedString)
			}
		})
	}
}

func convertToString(res Result) string {
	// var r string
	var r strings.Builder
	for key, val := range res.occurrences {
		r.WriteString(fmt.Sprintf("Word: '%s' at positions: ", key))
		for i := range val {
			r.WriteString(fmt.Sprintf("%d", val[i]))
			if i != len(val)-1 {
				// r = r + fmt.Sprintf(", ")
				r.WriteString(", ")
			}
		}
		// r = r + fmt.Sprintf(". ")
		r.WriteString(". ")
	}
	return r.String()
}
