package problem3

import "testing"

// Tests
func TestProblem3_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{
			name:  "Testcase 1 - input 13195",
			input: 13195,
			want:  29,
		},
		{
			name:  "Testcase 2 - input 600851475143",
			input: 600851475143,
			want:  6857,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem3(tt.input)

			if n != tt.want {
				t.Errorf("Problem3() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem3(600851475143)
	}
}
