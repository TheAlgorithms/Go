package charoccurrence

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string       // test description
	input    string       // user input
	expected map[rune]int // expected return
}{
	{
		"english text",
		"hello world!",
		map[rune]int{32: 1, 33: 1, 100: 1, 101: 1, 104: 1, 108: 3, 111: 2, 114: 1, 119: 1},
	},
	{
		"chinese text",
		" 世界 ",
		map[rune]int{32: 2, 19990: 1, 30028: 1},
	},
	{
		"persian text",
		"سلام دنیا!",
		map[rune]int{32: 1, 33: 1, 1575: 2, 1583: 1, 1587: 1, 1604: 1, 1605: 1, 1606: 1, 1740: 1},
	},
}

func TestCount(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := Count(test.input)
			eq := reflect.DeepEqual(test.expected, funcResult)
			if !eq {
				t.Errorf("Expected answer '%v' for string '%s' but answer given was %v", test.expected, test.input, funcResult)
			}
		})
	}
}
