package problem2

import "testing"

// Tests
func TestProblem2_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{
			name:  "Testcase 1 - input 10",
			input: 10,
			want:  10,
		},
		{
			name:  "Testcase 2 - input 100",
			input: 100,
			want:  44,
		},
		{
			name:  "Testcase 3 - input 4e6",
			input: 4e6,
			want:  4613732,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem2(tt.input)

			if n != tt.want {
				t.Errorf("Problem2() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem2(4e6)
	}
}
