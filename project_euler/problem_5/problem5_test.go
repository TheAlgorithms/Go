package problem5

import "testing"

// Tests
func TestProblem5_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{
			name:  "Testcase 1 - input 10",
			input: 10,
			want:  2520,
		},
		{
			name:  "Testcase 2 - input 20",
			input: 20,
			want:  232792560,
		},
		{
			name:  "Testcase 3 - input 5",
			input: 5,
			want:  60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem5(tt.input)

			if n != tt.want {
				t.Errorf("Problem5() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem5(20)
	}
}
