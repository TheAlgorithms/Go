package problem9

import "testing"

// Tests
func TestProblem9_Func(t *testing.T) {
	tests := []struct {
		name string
		want uint
	}{
		{
			name: "Testcase 1",
			want: 31875000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem9()

			if n != tt.want {
				t.Errorf("Problem9() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem9()
	}
}
