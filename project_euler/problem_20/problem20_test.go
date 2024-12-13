package problem20

import "testing"

// Tests
func TestProblem20_Func(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Problem 20 - Factorial digit sum", 10, 27},
		{"Problem 20 - Factorial digit sum", 100, 648},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem20(test.input)
			if got != test.expected {
				t.Errorf("Problem20() = got %v, want %v", got, test.expected)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem20_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem20(100)
	}
}
