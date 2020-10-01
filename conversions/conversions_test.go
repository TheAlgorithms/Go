package conversions

import "testing"

//BEGIN TESTS

func TestRomanToInteger(t *testing.T) {
	for _, test := range romanToIntegerTests{
		actual := romanToInteger(test.input)
		if actual != test.expected {
			t.Errorf("roman to integer test %s failed", test.name,)
		}
	}
}

//END TESTS

//BEGIN BENCHMARKS

func BenchmarkRomanToInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range romanToIntegerTests {
			romanToInteger(test.input)
		}
	}
}

//END BENCHMARKS