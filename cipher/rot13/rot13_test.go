package rot13

import (
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

func TestRot13Encrypt(t *testing.T) {
	for _, test := range rot13TestData {
		t.Run(test.description, func(t *testing.T) {
			input := test.input
			expected := test.expected
			assertRot13Output(t, input, expected)
		})
	}
}

func TestRot13Decrypt(t *testing.T) {
	for _, test := range rot13TestData {
		t.Run(test.description, func(t *testing.T) {
			input := test.expected
			expected := test.input
			assertRot13Output(t, input, expected)
		})
	}
}

func assertRot13Output(t *testing.T, input, expected string) {
	actual := rot13(input)
	if actual != expected {
		t.Fatalf("With input string %q was expecting %q but actual was %q",
			input, expected, actual)
	}
}

func FuzzRot13(f *testing.F) {
	for _, rot13TestInput := range rot13TestData {
		f.Add(rot13TestInput.input)
	}
	f.Fuzz(func(t *testing.T, input string) {
		if result := rot13(rot13(input)); result != input {
			t.Fatalf("With input string %q was expecting %q but actual was %q",
				input, input, result)
		}
	})
}
