package strings_test

import (
	"reflect"
	"testing"

	"github.com/TheAlgorithms/Go/strings"
)

func TestIsSubsequence(t *testing.T) {
	var testCases = []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			"Valid case 1 ",
			"ace",
			"abcde",
			true,
		},
		{
			"Invalid case 1",
			"aec",
			"abcde",
			false,
		},
		{
			"Empty strings",
			"",
			"",
			true,
		},
		{
			"s is more then t",
			"aeccccc",
			"abcde",
			false,
		},
		{
			"s is empty",
			"",
			"abcde",
			true,
		},
		{
			"Equal strings",
			"aec",
			"aec",
			true,
		},
		{
			"Valid case 2",
			"pyr",
			"wpxqyrz",
			true,
		},
		{
			"Invalid case 2",
			"prx",
			"wpxqyrz",
			false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := strings.IsSubsequence(test.s, test.t)
			if !reflect.DeepEqual(test.expected, funcResult) {
				t.Errorf("expected: %v, got %v", test.expected, funcResult)
			}
		})
	}
}
