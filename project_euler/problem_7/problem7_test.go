package problem7

import "testing"

// Tests
func TestProblem7_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  int64
	}{
		{
			name:  "Testcase 1 - input 6",
			input: 6,
			want:  13,
		},
		{
			name:  "Testcase 2 - input 10001",
			input: 10001,
			want:  104743,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem7(tt.input)

			if n != tt.want {
				t.Errorf("Problem7() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem7(10001)
	}
}
