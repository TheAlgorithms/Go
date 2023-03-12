package strings_test

import (
	"github.com/TheAlgorithms/Go/strings"
	"reflect"
	"testing"
)

func TestCountChars(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected map[rune]int
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
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := strings.CountChars(test.input)
			if !reflect.DeepEqual(test.expected, funcResult) {
				t.Errorf("expected: %v, got %v", test.expected, funcResult)
			}
		})
	}
}
