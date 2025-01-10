package problem6

import "testing"

// Tests
func TestProblem6_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{
			name:  "Testcase 1 - input 10",
			input: 10,
			want:  2640,
		},
		{
			name:  "Testcase 2 - input 100",
			input: 100,
			want:  25164150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem6(tt.input)

			if n != tt.want {
				t.Errorf("Problem6() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem6(100)
	}
}
