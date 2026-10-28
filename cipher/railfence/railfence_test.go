package railfence

import (
	"testing"
	"unicode/utf8"
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
		{
			"Encrypt with Unicode characters",
			"Hello, 世界!",
			2,
			"Hlo 界el,世!",
		},
		{
			"Encrypt with edge case: 0 rails",
			"hello",
			0,
			"hello",
		},
		{
			"Encrypt with edge case: negative rails",
			"hello",
			-5,
			"hello",
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
		{
			"Decrypt with Unicode characters",
			"Hlo 界el,世!",
			2,
			"Hello, 世界!",
		},
		{
			"Decrypt with edge case: 0 rails",
			"hello",
			0,
			"hello",
		},
		{
			"Decrypt with edge case: negative rails",
			"hello",
			-5,
			"hello",
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

func FuzzRailfence(f *testing.F) {
	testCases := []struct {
		input string
		rails int
	}{
		{"hello", 2},
		{"hello world", 3},
		{"hello", 1},
		{"hi", 100},
		{"Hello, 世界!", 2},
		{"The Quick Brown Fox Jumps Over The Lazy Dog", 4},
	}

	for _, tc := range testCases {
		f.Add(tc.input, tc.rails)
	}

	f.Fuzz(func(t *testing.T, input string, rails int) {
		if !utf8.ValidString(input) {
			return
		}
		if rails <= 0 {
			if Encrypt(input, rails) != input {
				t.Errorf("Encrypt with rails <= 0 should return input unmodified")
			}
			if Decrypt(input, rails) != input {
				t.Errorf("Decrypt with rails <= 0 should return input unmodified")
			}
			return
		}
		encrypted := Encrypt(input, rails)
		decrypted := Decrypt(encrypted, rails)
		if decrypted != input {
			t.Fatalf("With input: %q and rails: %d\n\tExpected: %q\n\tGot: %q", input, rails, input, decrypted)
		}
	})
}
