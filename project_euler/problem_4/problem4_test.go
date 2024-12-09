package problem4

import "testing"

// Tests
func TestProblem4_Func(t *testing.T) {
	tests := []struct {
		name string
		want uint
	}{
		{
			name: "Testcase 1",
			want: 906609,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem4()

			if n != tt.want {
				t.Errorf("Problem4() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem4()
	}
}
