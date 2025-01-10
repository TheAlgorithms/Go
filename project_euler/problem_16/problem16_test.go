package problem16

import "testing"

// Tests
func TestProblem16_Func(t *testing.T) {
	tests := []struct {
		name     string
		exponent int64
		want     int64
	}{
		{"2^15", 15, 26},
		{"2^1000", 1000, 1366},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem16(tt.exponent); got != tt.want {
				t.Errorf("Problem16() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark
func BenchmarkProblem16_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem16(1000)
	}
}
