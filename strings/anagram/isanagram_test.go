package anagram

import (
	"testing"
)

var testCases = []struct {
	name     string // test description
	input1   string // user input1
	input2   string //user input 2
	expected bool   // expected return
}{
	{
		"Anagram String 1",
		"angel",
		"glean",
		true,
	},
	{
		"Non-Anagram String 1",
		"git",
		"hub",
		false,
	},
	{
		"Anagram String 2",
		"mango",
		"goman",
		true,
	},
	{
		"Non-Anagram String 2",
		"length",
		"different",
		false,
	},
	{
		"Anagram String 3",
		"distill231",
		"23sidtill1",
		true,
	},
	{
		"Anagram String 4",
		"I am happy",
		"Happy am I",
		true,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			func_result := isanagram(test.input1, test.input2)
			if test.expected != func_result {
				t.Errorf("Expected answer '%t' for string '%s' and '%s' but answer given was %t", test.expected, test.input1, test.input2, func_result)
			}
		})
	}
}
