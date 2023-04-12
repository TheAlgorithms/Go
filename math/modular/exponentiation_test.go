// exponentiation_test.go
// description: Test for ModularExponentiation
// author(s) [Taj](https://github.com/tjgurwara99)
// see exponentiation.go

package modular

import "testing"

type cases struct {
	name          string
	description   string
	base          int64
	exponent      int64
	mod           int64
	expected      int64
	expectedError error
}

var testCases = []cases{
	{
		name:          "Test 1",
		description:   "Test 1: 3^6 % 3 == 0",
		base:          3,
		exponent:      6,
		mod:           3,
		expected:      0,
		expectedError: nil,
	},
	{
		name:          "Test 2",
		description:   "Test 2: 33^60 % 25 == 1",
		base:          33,
		exponent:      60,
		mod:           25,
		expected:      1,
		expectedError: nil,
	},
	{
		name:          "Test 3",
		description:   "Test 3: 17^60 % 23 == 2",
		base:          17,
		exponent:      60,
		mod:           23,
		expected:      2,
		expectedError: nil,
	},
	{
		name:          "Test 4",
		description:   "Test 4: 17^60 % 1 == 0", // handling result when we get mod = 1
		base:          17,
		exponent:      60,
		mod:           1,
		expected:      0,
		expectedError: nil,
	},
	{
		name:          "Error test 1",
		description:   "Testing whether we receive the expected errors gracefully",
		base:          50,
		exponent:      -1,
		mod:           2,
		expected:      -1,
		expectedError: ErrorNegativeExponent,
	},
}

func TestExponentiation(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := Exponentiation(test.base, test.exponent, test.mod)
			if err != test.expectedError {
				t.Logf("Test Failed for %s", test.name)
				t.Logf("Unexpected error occurred")
				t.Errorf("Expected error: %v, Received error: %v", test.expectedError, err)
			}
			if result != test.expected {
				t.Logf("Test Failed for %s", test.description)
				t.Fatalf("Expected: %d, Received: %d", test.expected, result)
			}
		})
	}
}

func BenchmarkExponentiation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Exponentiation(17, 60, 23)
	}
}
