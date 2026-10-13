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
			"Encrypt with Unicode runes",
			"café crème",
			3,
			"c maécèefr",
		},
		{
			"Encrypt with negative rails",
			"hello",
			-1,
			"hello",
		},
		{
			"Encrypt with zero rails",
			"hello",
			0,
			"hello",
		},
		{
			"Encrypt with empty string",
			"",
			3,
			"",
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
			"Decrypt with Unicode runes",
			"c maécèefr",
			3,
			"café crème",
		},
		{
			"Decrypt with negative rails",
			"hello",
			-1,
			"hello",
		},
		{
			"Decrypt with zero rails",
			"hello",
			0,
			"hello",
		},
		{
			"Decrypt with empty string",
			"",
			3,
			"",
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

func FuzzRailFence(f *testing.F) {
	seeds := []struct {
		input string
		rails int
	}{
		{"hello", 2},
		{"hello world", 3},
		{"hello", 1},
		{"hi", 100},
		{"café crème", 3},
		{"the quick brown fox jumps over the lazy dog", 4},
		{"1234567890", 2},
		{"", 3},
		{"a", 2},
		{"hello", -1},
		{"hello", 0},
	}

	for _, seed := range seeds {
		f.Add(seed.input, seed.rails)
	}

	f.Fuzz(func(t *testing.T, input string, rails int) {
		if !utf8.ValidString(input) {
			return
		}
		encrypted := Encrypt(input, rails)
		decrypted := Decrypt(encrypted, rails)
		if decrypted != input {
			t.Fatalf("round-trip failed: original=%q, rails=%d, encrypted=%q, decrypted=%q", input, rails, encrypted, decrypted)
		}
	})
}
