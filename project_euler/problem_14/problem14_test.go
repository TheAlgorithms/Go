package problem14

import "testing"

// Tests
func TestProblem14_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint64
	}{
		{"Input 30", 30, 27},
		{"Input 1e6", 1e6, 837799},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Problem14(tt.input)
			if got != tt.want {
				t.Errorf("Problem14() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem14_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem14(1e6)
	}
}
