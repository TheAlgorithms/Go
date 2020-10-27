package power

import "testing"

var testCases = []struct {
	name     string
	base     uint
	power    uint
	expected uint
}{
	{"0^2", 0, 2, 0},
	{"2^0", 2, 0, 1},
	{"2^3", 2, 3, 8},
	{"8^3", 8, 3, 512},
	{"10^5", 10, 5, 100000},
}

func TestIterativePower(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IterativePower(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
}

func TestRecursivePower(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RecursivePower(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
}

func TestRecursivePower1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RecursivePower1(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
}

func BenchmarkIterativePower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativePower(10, 5)
	}
}

func BenchmarkRecursivePower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursivePower(10, 5)
	}
}

func BenchmarkRecursivePower1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursivePower1(10, 5)
	}
}
