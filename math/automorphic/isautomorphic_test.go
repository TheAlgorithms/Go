// isautomorphic_test.go

// Package sumofdigits provides method to determine whether a number is automorphic or not.
package automorphic

// testing package: to easily perform testing
import (
	"testing"
)

var testCases = []struct {
	name     string // test description
	input    int    // user input
	expected bool   // expected return
}{
	{
		"negetive number: not Automorphic",
		-1,
		false,
	},
	{
		"negetive number: not Automorphic",
		-146,
		false,
	},
	{
		"0: is Automorphic",
		0,
		true,
	},
	{
		"1: is Automorphic",
		1,
		true,
	},
	{
		"7: not Automorphic",
		7,
		false,
	},
	{
		"83: not Automorphic",
		83,
		false,
	},
	{
		"376: is Automorphic",
		376,
		true,
	},
}

func TestIsAutomorphic(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := IsAutomorphic(test.input)
			if test.expected != funcResult {
				t.Errorf("Expected answer '%t' for the number '%d' but answer given was %t", test.expected, test.input, funcResult)
			}
		})
	}
}
