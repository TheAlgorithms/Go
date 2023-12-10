package parenthesis

import (
	"testing"
)

var parenthesisTestCases = []struct {
	name     string
	text     string
	expected bool
}{
	{
		"simple test with one level deep",
		"(3*9-2)+(3/7^3)",
		true,
	},
	{
		"three nested parenthesis with three deep level",
		"(-1*(5+2^(3-4)*(7.44-12)+6.66/(3.43-(1+2)))*(sqrt(3-4)))",
		true,
	},
	{
		"one opened parenthesis without be closed",
		"(2*9-17)*((7+3/3)*2*(-1+4)",
		false,
	},
	{
		"one open parenthesis for each close one but not in pairs",
		"(4*(39.22-7.4)/6.77))(",
		false,
	},
	{
		"6 deep level",
		"(5+33.2)*((8.04-6.5)*(1 - (1^(2-1.22*5.44+(4.33*7.2^(0.34*(-1.23+2)))))))",
		true,
	},
	{
		"inverted parenthesis",
		")()()()()(",
		false,
	},
}

func TestParenthesis(t *testing.T) {
	for _, tc := range parenthesisTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Parenthesis(tc.text)
			if actual != tc.expected {
				t.Errorf("Expected %t value from test %s with input %s\n", tc.expected, tc.name, tc.text)
			}
		})
	}
}
