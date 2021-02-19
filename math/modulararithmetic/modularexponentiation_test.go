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
	for i := 0; i < b.N; i++ {
		_ = ModularExponentiation(17, 60, 23)
	}
}
