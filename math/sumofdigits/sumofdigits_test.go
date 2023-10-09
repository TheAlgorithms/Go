// sumofdigits_test.go

package sumofdigits

// testing package: to easily perform testing
import (
	"testing"
)

var testCases = []struct {
	name     string // test description
	input    int    // user input
	expected int    // expected return
}{
	{
		"negetive number: returning -1",
		-1,
		-1,
	},
	{
		"negetive number: returning -1",
		-146,
		-1,
	},
	{
		"large negetive number: returning -1",
		-787485,
		-1,
	},
	{
		"0: should give 0",
		0,
		0,
	},
	{
		"single digit positive number: should give number itself",
		4,
		4,
	},
	{
		"medium positive number: should give digit sum",
		606054,
		21,
	},
	{
		"medium positive number: should give digit sum",
		123451,
		16,
	},
	{
		"large positive number: should give digit sum",
		75539676,
		48,
	},
	{
		"large positive number: should give digit sum",
		857353712,
		41,
	},
}

func TestSumofdigits(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			funcResult := Sumofdigits(test.input)
			if test.expected != funcResult {
				t.Errorf("Expected answer '%d' for the number '%d' but answer given was %d", test.expected, test.input, funcResult)
			}
		})
	}
}
