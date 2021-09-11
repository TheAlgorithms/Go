package gcd

import "testing"

type testFunction func(int64, int64) int64

var testCases = []struct {
	name   string
	a      int64
	b      int64
	output int64
}{
	{"gcd of 10 and 0", 10, 0, 10},
	{"gcd of 98 and 56", 98, 56, 14},
	{"gcd of 0 and 10", 0, 10, 10},
}

func TemplateTestGCD(t *testing.T, f testFunction) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := f(tc.a, tc.b)
			if actual != tc.output {
				t.Errorf("Expected GCD of %d and %d to be: %v, but got: %d", tc.a, tc.b, tc.output, actual)
			}
		})
	}
}

func TestGCDRecursive(t *testing.T) {
	TemplateTestGCD(t, Recursive)
}

func TestGCDIterative(t *testing.T) {
	TemplateTestGCD(t, Iterative)
}

func TemplateBenchmarkGCD(b *testing.B, f testFunction) {
	for i := 0; i < b.N; i++ {
		f(98, 56)
	}
}

func BenchmarkGCDRecursive(b *testing.B) {
	TemplateBenchmarkGCD(b, Recursive)
}

func BenchmarkGCDIterative(b *testing.B) {
	TemplateBenchmarkGCD(b, Iterative)
}
