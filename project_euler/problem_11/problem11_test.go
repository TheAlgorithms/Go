package problem11

import "testing"

// Tests
func TestProblem11_Func(t *testing.T) {
	testCases := []struct {
		name     string
		expected uint
	}{
		{"Test Case 1", 70600674},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Problem11()
			if actual != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, actual)
			}
		})
	}
}

// Benchmark
func BenchmarkProblem11_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem11()
	}
}
