package ceaser

import (
	"testing"
)

var caesarTestData = []struct {
	description string
	input       string
	key         int
	expected    string
}{
	{
		"Basic caesar encryption with letter 'a'",
		"a",
		3,
		"d",
	},
	{
		"Basic caesar encryption wrap around alphabet on letter 'z'",
		"z",
		3,
		"c",
	},
	{
		"Encrypt a simple string with ceasar encryiption",
		"hello",
		3,
		"khoor",
	},
	{
		"Encrypt a simple string with key 13",
		"hello",
		13,
		"uryyb",
	},
	{
		"Encrypt a simple string with key -13",
		"hello",
		-13,
		"uryyb",
	},
	{
		"With key of 26 output should be the same as the input",
		"no change",
		26,
		"no change",
	},
	{
		"Encrpyt sentence with key 10",
		"the quick brown fox jumps over the lazy dog.",
		10,
		"dro aesmu lbygx pyh tewzc yfob dro vkji nyq.",
	},
}

func TestCeasarCipher(t *testing.T) {
	for _, test := range caesarTestData {
		t.Run(test.description, func(t *testing.T) {
			actual := caesarCipher(test.input, test.key)
			if actual != test.expected {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("With input string '%s' and key '%d' was expecting '%s' but actual was '%s'",
					test.input, test.key, test.expected, actual)
			}
		})
	}
}
