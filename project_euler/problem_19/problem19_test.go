package problem19

import "testing"

// Tests
func TestProblem19_Func(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{"Problem 19 - Counting Sundays", 171},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem19()
			if got != test.expected {
				t.Errorf("Problem19() = got %v, want %v", got, test.expected)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem19_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem19()
	}
}
