package strings_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/strings"
)

var testCases = []struct {
	name        string
	input       string
	order       strings.IsogramOrder
	expectedVal bool
	expectedErr error
}{
	{
		"Alphanumeric string 1",
		"copy1",
		1,
		false,
		errors.New("Cannot contain numbers or symbols"),
	},
	{
		"Alphanumeric string 2",
		"copy1sentence",
		1,
		false,
		errors.New("Cannot contain numbers or symbols"),
	},
	{
		"Alphanumeric string 3",
		"copy1 sentence with space",
		1,
		false,
		errors.New("Cannot contain numbers or symbols"),
	},
	{
		"Alphabetic string 1",
		"allowance",
		1,
		false,
		nil,
	},
	{
		"Alphabetic string 2",
		"My Doodle",
		1,
		false,
		nil,
	},
	{
		"Alphabetic string with symbol",
		"Isogram!",
		1,
		false,
		errors.New("Cannot contain numbers or symbols"),
	},
	{
		"Isogram string 1",
		"Uncopyrightable",
		1,
		true,
		nil,
	},
	{
		"Second order isogram 1",
		"Caucasus",
		2,
		true,
		nil,
	},
	{
		"Second order isogram 2",
		"Couscous",
		2,
		true,
		nil,
	},
	{
		"Third order isogram 1",
		"Deeded",
		3,
		true,
		nil,
	},
	{
		"Third order isogram 2",
		"Sestettes",
		3,
		true,
		nil,
	},
	{
		"Not an isogram",
		"Pneumonoultramicroscopicsilicovolcanoconiosis",
		1,
		false,
		nil,
	},
	{
		"Not an isogram",
		"Randomstring",
		4,
		false,
		errors.New("Invalid isogram order provided"),
	},
	{
		"Third order isogram checked as first order",
		"Deeded",
		1,
		false,
		nil,
	},
}

func TestIsIsogram(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualVal, actualErr := strings.IsIsogram(test.input, test.order)

			if actualErr != nil && test.expectedErr.Error() != actualErr.Error() {
				t.Errorf("Expected to be '%v' for string %s but was '%v'.", test.expectedErr, test.input, actualErr)
			}

			if test.expectedVal != actualVal {
				t.Errorf("Expected to be '%v' for string %s but was '%v'.", test.expectedVal, test.input, actualVal)
			}
		})
	}
}
