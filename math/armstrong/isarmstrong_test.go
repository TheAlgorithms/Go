// isarmstrong_test.go

package armstrong

import "testing"

var testCases = []struct {
	name     string // test description
	input    int    // user input
	expected bool   // expected return
}{
	{
		"negative number: Not an armstrong number",
		-140,
		false,
	},
	{
		"positive number: Not an armstrong number",
		23,
		false,
	},
	{
		"smallest armstrong number",
		0,
		true,
	},
	{
		"smallest armstrong number with more than one digit",
		153,
		true,
	},
	{
		"random armstrong number",
		407,
		true,
	},
	{
		"random armstrong number",
		9474,
		true,
	},
}

func TestIsArmstrong(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := IsArmstrong(test.input)
			if test.expected != funcResult {
				t.Errorf("Expected answer '%t' for the number '%d' but answer given was %t", test.expected, test.input, funcResult)
			}
		})
	}
}
