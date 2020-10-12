package gcd

import "testing"

var testCases = []struct {
	name   string
	a      int
	b      int
	output int
}{
	{"gcd of 10 and 0", 10, 0, 10},
	{"gcd of 98 and 56", 98, 56, 14},
	{"gcd of 0 and 10", 0, 10, 10},
}

func TestGCD(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Gcd(tc.a, tc.b)
			if actual != tc.output {
				t.Errorf("Expected GCD of %d and %d to be: %v, but got: %d", tc.a, tc.b, tc.output, actual)
			}
		})
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gcd(98, 56)
	}
}
