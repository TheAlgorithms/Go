package problem1

import "testing"

// Tests
func TestProblem1_Func(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint
		want      uint
	}{
		{
			name:      "Testcase 1 - threshold 10",
			threshold: 10,
			want:      23,
		},
		{
			name:      "Testcase 2 - threshold 100",
			threshold: 100,
			want:      2318,
		},
		{
			name:      "Testcase 3 - threshold 1000",
			threshold: 1000,
			want:      233168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem1(tt.threshold)

			if n != tt.want {
				t.Errorf("Problem1() = %v, want %v", n, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Problem1(1000)
	}
}
