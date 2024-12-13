package problem17

import "testing"

// Tests
func TestProblem17_Func(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"1 to 5", "one two three four five", 19},
		{"1 to 1000", INPUT, 21124},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem17(tt.input); got != tt.want {
				t.Errorf("Problem17() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark
func BenchmarkProblem17_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem17(INPUT)
	}
}
