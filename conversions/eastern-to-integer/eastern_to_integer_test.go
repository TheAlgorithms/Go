package easterntointeger

import "testing"

type easternNumberToIntegerConversionTest struct {
	input    string
	expected int
	name     string
}

var easternToTests = []easternNumberToIntegerConversionTest{
	{input: "۵۲۸۹", expected: 5289, name: "5289"},
	{input: "۸٨٤٥", expected: 8845, name: "8845"},
	{input: "٥۵۵", expected: 555, name: "555-mixed"},
	{input: "", expected: 0, name: "empty"},
	{input: "۰", expected: 0, name: "zero"},
	{input: "۰۲۹", expected: 29, name: "029"},
	{input: "abcd", expected: 0, name: "words"},
}

func TestEasternToInteger(t *testing.T) {

	for _, test := range easternToTests {
		convertedValue := EasternToInteger(test.input)
		if convertedValue != test.expected {
			t.Errorf(
				"eastern to integer test %s failed. expected '%d' but got '%d'",
				test.name, test.expected, convertedValue,
			)
		}
	}
}
