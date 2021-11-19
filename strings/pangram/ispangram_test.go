package pangram

import (
	"testing"
)

var testCases = []struct {
	name     string // test description
	input    string // user input
	expected bool   // expected return
}{
	{
		"empty string",
		"",
		false,
	},
	{
		"non pangram string without spaces",
		"abc",
		false,
	},
	{
		"non pangram string with spaces",
		"Hello World",
		false,
	},
	{
		"Pangram string 1",
		" Abcdefghijklmnopqrstuvwxyz",
		true,
	},
	{
		"pangram string 2",
		"cdefghijklmnopqrstuvwxABC zyb",
		true,
	},
	{
		"pangram string 3",
		"The Quick Brown Fox Jumps Over the Lazy Dog",
		true,
	},
}

func TestIsPangram(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			func_result := IsPangram(test.input)
			if test.expected != func_result {
				t.Errorf("Expected answer '%t' for string '%s' but answer given was %t", test.expected, test.input, func_result)
			}
		})
	}
}
