// a1z26_test.go
// description: Implementation of test cases for the a1z26 cipher
// details:
// The A1Z26 cipher is a simple substiution cipher where each letter is
// replaced by the number of the order they're in. For example, A corresponds to
// 1, B = 2, C = 3, etc..
//
// author(s): [Focusucof](https://github.com/Focusucof)
// see a1z26.go
package a1z26

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	var a1z26TestData = []struct {
		description string
		input       string
		expected    string
	}{
		{
			"A1z26 encryption with the letter 'a'",
			"a",
			"1",
		},
		{
			"A1z26 encryption with the letter 'z'",
			"z",
			"26",
		},
		{
			"A1z26 encryption with the string \"Hello World\"",
			"Hello World",
			"8-5-12-12-15 23-15-18-12-4",
		},
		{
			"A1z26 encryption with the string \"HELLO WORLD\" in all caps",
			"HELLO WORLD",
			"8-5-12-12-15 23-15-18-12-4",
		},
		{
			"A1z26 encryption with the string \"The quick brown fox jumps over the lazy dog\"",
			"The quick brown fox jumps over the lazy dog",
			"20-8-5 17-21-9-3-11 2-18-15-23-14 6-15-24 10-21-13-16-19 15-22-5-18 20-8-5 12-1-26-25 4-15-7",
		},
	}
	for _, test := range a1z26TestData {
		t.Run(test.description, func(t *testing.T) {
			actual := Encrypt(test.input)
			if actual != test.expected {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("With input string '%s' was expecting '%s' but actual was '%s'",
					test.input, test.expected, actual)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	var a1z26TestData = []struct {
		description string
		input       string
		expected    string
	}{
		{
			"A1z26 decryption with the letter '1'",
			"1",
			"a",
		},
		{
			"A1z26 decryption with the letter '26'",
			"26",
			"z",
		},
		{
			"A1z26 decryption with the string \"8-5-12-12-15 23-15-18-12-4\"",
			"8-5-12-12-15 23-15-18-12-4",
			"hello world",
		},
		{
			"A1z26 encryption with the string \"20-8-5 17-21-9-3-11 2-18-15-23-14 6-15-24 10-21-13-16-19 15-22-5-18 20-8-5 12-1-26-25 4-15-7\"",
			"20-8-5 17-21-9-3-11 2-18-15-23-14 6-15-24 10-21-13-16-19 15-22-5-18 20-8-5 12-1-26-25 4-15-7",
			"the quick brown fox jumps over the lazy dog",
		},
	}
	for _, test := range a1z26TestData {
		t.Run(test.description, func(t *testing.T) {
			actual := Decrypt(test.input)
			if actual != test.expected {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("With input string '%s' was expecting '%s' but actual was '%s'",
					test.input, test.expected, actual)
			}
		})
	}
}
