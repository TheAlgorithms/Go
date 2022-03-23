package cipher_test

import (
	"github.com/TheAlgorithms/Go/cipher"
	"testing"
)

var rot13TestData = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Basic rotation with letter 'a' gives 'n",
		"a",
		"n",
	},
	{
		"Rotation with wrapping around alphabet on letter 'z' gives 'm'",
		"z",
		"m",
	},
	{
		"Rotation on 'hello world'",
		"hello world",
		"uryyb jbeyq",
	},
	{
		"Rotation on the rotation of 'hello world' gives 'hello world' back",
		"uryyb jbeyq",
		"hello world",
	},
	{
		"Full sentence rotation",
		"the quick brown fox jumps over the lazy dog.",
		"gur dhvpx oebja sbk whzcf bire gur ynml qbt.",
	},
	{
		"Sentence from Rot13.go main function",
		"we'll just make him an offer he can't refuse... tell me you get the pop culture reference",
		"jr'yy whfg znxr uvz na bssre ur pna'g ershfr... gryy zr lbh trg gur cbc phygher ersrerapr",
	},
}

func TestRot13(t *testing.T) {
	for _, test := range rot13TestData {
		t.Run(test.description, func(t *testing.T) {
			input := test.input
			expected := test.expected
			assertRot13Output(t, input, expected)
		})
	}
}

func assertRot13Output(t *testing.T, input, expected string) {
	actual, err := cipher.Rot13(input)
	if err != nil {
		t.Fatalf("FAIL: no errors expected: %s", err)
	}
	if actual != expected {
		t.Fatalf("With input string '%s' was expecting '%s' but actual was '%s'",
			input, expected, actual)
	}
}
