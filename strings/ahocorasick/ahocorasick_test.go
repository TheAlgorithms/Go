package ahocorasick

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

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
	var r strings.Builder
	for key, val := range res.occurrences {
		r.WriteString(fmt.Sprintf("Word: '%s' at positions: ", key))
		for i := range val {
			r.WriteString(fmt.Sprintf("%d", val[i]))
			if i != len(val)-1 {
				r.WriteString(", ")
			}
		}
		r.WriteString(". ")
	}
	return r.String()
}
