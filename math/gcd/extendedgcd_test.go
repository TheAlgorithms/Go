package gcd

import "testing"

type testExtendedFunction func(int64, int64) (int64, int64, int64)

func TemplateTestExtendedGCD(t *testing.T, f testExtendedFunction) {
	var testCasesExtended = []struct {
		name string
		a    int64
		b    int64
		gcd  int64
		x    int64
		y    int64
	}{
		{"gcd of 10 and 0", 10, 0, 10, 1, 0},
		{"gcd of 98 and 56", 98, 56, 14, -1, 2},
		{"gcd of 0 and 10", 0, 10, 10, 0, 1},
	}
	for _, tc := range testCasesExtended {
		t.Run(tc.name, func(t *testing.T) {
			actualGcd, actualX, actualY := f(tc.a, tc.b)
			if actualGcd != tc.gcd {
				t.Errorf("Expected GCD of %d and %d to be: %v, but got: %d", tc.a, tc.b, tc.gcd, actualGcd)
			}
			if actualX != tc.x {
				t.Errorf("Expected x satisfying %d * x + %d * y = gcd to be: %v, but got: %d", tc.a, tc.b, tc.x, actualX)
			}
			if actualY != tc.y {
				t.Errorf("Expected y satisfying %d * x + %d * y = gcd to be: %v, but got: %d", tc.a, tc.b, tc.y, actualY)
			}
		})
	}
}

func TestExtendedGCDRecursive(t *testing.T) {
	TemplateTestExtendedGCD(t, ExtendedRecursive)
}

func TestExtendedGCDIterative(t *testing.T) {
	TemplateTestExtendedGCD(t, ExtendedIterative)
}

func TemplateBenchmarkExtendedGCD(b *testing.B, f testExtendedFunction) {
	for i := 0; i < b.N; i++ {
		f(98, 56)
	}
}

func BenchmarkExtendedGCDRecursive(b *testing.B) {
	TemplateBenchmarkExtendedGCD(b, ExtendedRecursive)
}

func BenchmarkExtendedGCDIterative(b *testing.B) {
	TemplateBenchmarkExtendedGCD(b, ExtendedIterative)
}
