package problem8

import "testing"

// Tests
func TestProblem8_Func(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint
	}{
		{
			name:  "Testcase 1 - input 4",
			input: 4,
			want:  5832,
		},
		{
			name:  "Testcase 2 - input 13",
			input: 13,
			want:  23514624000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem8(tt.input)

			if n != tt.want {
				t.Errorf("Problem8() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem8(13)
	}
}
