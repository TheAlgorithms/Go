package problem10

import "testing"

// Tests
func TestProblem10_Func(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  uint
	}{
		{
			name:  "Testcase 1 - input 10",
			input: 10,
			want:  17,
		},
		{
			name:  "Testcase 2 - input 2000000",
			input: 2000000,
			want:  142913828922,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem10(tt.input)

			if n != tt.want {
				t.Errorf("Problem10() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem10(2000000)
	}
}
