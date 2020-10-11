package main

import (
	"testing"
)

var diffieHellmanTestData = []struct {
	description string
	b, e, mod   int
	expected    int
}{
	{
		"Basic exponentiation with base 1, exponent 1 and modulo 1",
		1, 1, 1,
		0,
	},
	{
		"Basic exponentiation with base 2, exponent 10 and modulo 5",
		2, 10, 5,
		4,
	},
	{
		"Exponentiation test with base 3, exponent 3 and modulo 2",
		3, 3, 2,
		1,
	},
	{
		"Exponentiation test with base 5, exponent 3 and modulo 5",
		5, 3, 5,
		0,
	},
	{
		"Exponentiation test with base 10, exponent 2 and modulo 12",
		10, 2, 12,
		4,
	},
	{
		"Exponentiation test with base 15, exponent 20 and modulo 21",
		15, 20, 21,
		15,
	},
}

func TestModularExponentiation(t *testing.T) {
	for _, test := range diffieHellmanTestData {
		t.Run(test.description, func(t *testing.T) {
			actual := modularExponentiation(test.b, test.e, test.mod)
			if actual != test.expected {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %d, actual %d", test.expected, actual)
			}
		})
	}
}
