package modulararithmetic

import "testing"

type cases struct {
	name        string
	description string
	base        int64
	exponent    int64
	mod         int64
	expected    int64
}

var testCases = []cases{
	{
		name:        "Test 1",
		description: "Test 1: 3^6 % 3 == 0",
		base:        3,
		exponent:    6,
		mod:         3,
		expected:    0,
	},
	{
		name:        "Test 2",
		description: "Test 2: 33^60 % 25 == 1",
		base:        33,
		exponent:    60,
		mod:         25,
		expected:    1,
	},
	{
		name:        "Test 3",
		description: "Test 3: 17^60 % 23 == 2",
		base:        17,
		exponent:    60,
		mod:         23,
		expected:    2,
	},
	{
		name:        "Test 4",
		description: "Test 4: 3^6 % 3 == 0",
		base:        3,
		exponent:    6,
		mod:         3,
		expected:    0,
	},
	{
		name:        "Test 5",
		description: "Test 5: 33^60 % 25 == 1",
		base:        33,
		exponent:    60,
		mod:         25,
		expected:    1,
	},
	{
		name:        "Test 6",
		description: "Test 6: 17^60 % 23 == 2",
		base:        17,
		exponent:    60,
		mod:         23,
		expected:    2,
	},
	{
		name:        "Test 7",
		description: "Test 7: 3^6 % 3 == 0",
		base:        3,
		exponent:    6,
		mod:         3,
		expected:    0,
	},
	{
		name:        "Test 8",
		description: "Test 8: 33^60 % 25 == 1",
		base:        33,
		exponent:    60,
		mod:         25,
		expected:    1,
	},
	{
		name:        "Test 9",
		description: "Test 9: 17^60 % 23 == 2",
		base:        17,
		exponent:    60,
		mod:         23,
		expected:    2,
	},
	{
		name:        "Test 10",
		description: "Test 10: 3^6 % 3 == 0",
		base:        3,
		exponent:    6,
		mod:         3,
		expected:    0,
	},
	{
		name:        "Test 10",
		description: "Test 10: 33^60 % 25 == 1",
		base:        33,
		exponent:    60,
		mod:         25,
		expected:    1,
	},
	{
		name:        "Test 11",
		description: "Test 11: 17^60 % 23 == 2",
		base:        17,
		exponent:    60,
		mod:         23,
		expected:    2,
	},
	{
		name:        "Test 12",
		description: "Test 12: 3^6 % 3 == 0",
		base:        3,
		exponent:    6,
		mod:         3,
		expected:    0,
	},
	{
		name:        "Test 13",
		description: "Test 13: 33^60 % 25 == 1",
		base:        33,
		exponent:    60,
		mod:         25,
		expected:    1,
	},
	{
		name:        "Test 14",
		description: "Test 14: 17^60 % 23 == 2",
		base:        17,
		exponent:    60,
		mod:         23,
		expected:    2,
	},
}

func TestModularExponentation(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := ModularExponentiation(test.base, test.exponent, test.mod)
			if result != test.expected {
				t.Logf("Test Failed for %s", test.description)
				t.Fatalf("Expected: %d, Recieved: %d", test.expected, result)
			}
		})
	}
}

func BenchmarkModularExponentiation(b *testing.B) {
	for _, test := range testCases {
		_ = ModularExponentiation(test.base, test.exponent, test.mod)
	}
}
