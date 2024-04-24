package railfence

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	var railFenceTestData = []struct {
		description string
		input       string
		rails       int
		expected    string
	}{
		{
			"Encrypt with 2 rails",
			"hello",
			2,
			"hloel",
		},
		{
			"Encrypt with 3 rails",
			"hello world",
			3,
			"horel ollwd",
		},
		{
			"Encrypt with edge case: 1 rail",
			"hello",
			1,
			"hello",
		},
		{
			"Encrypt with more rails than letters",
			"hi",
			100,
			"hi",
		},
	}

	for _, test := range railFenceTestData {
		t.Run(test.description, func(t *testing.T) {
			actual := Encrypt(test.input, test.rails)
			if actual != test.expected {
				t.Errorf("FAIL: %s - Encrypt(%s, %d) = %s, want %s", test.description, test.input, test.rails, actual, test.expected)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	var railFenceTestData = []struct {
		description string
		input       string
		rails       int
		expected    string
	}{
		{
			"Decrypt with 2 rails",
			"hloel",
			2,
			"hello",
		},
		{
			"Decrypt with 3 rails",
			"ho l lewrdlo",
			3,
			"hld olle wor",
		},
		{
			"Decrypt with edge case: 1 rail",
			"hello",
			1,
			"hello",
		},
		{
			"Decrypt with more rails than letters",
			"hi",
			100,
			"hi",
		},
	}

	for _, test := range railFenceTestData {
		t.Run(test.description, func(t *testing.T) {
			actual := Decrypt(test.input, test.rails)
			if actual != test.expected {
				t.Errorf("FAIL: %s - Decrypt(%s, %d) = %s, want %s", test.description, test.input, test.rails, actual, test.expected)
			}
		})
	}
}
