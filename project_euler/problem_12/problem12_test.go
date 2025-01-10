package problem12

import "testing"

func TestProblem12_Func(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{"Test Case 1", 6, 28},
		{"Test Case 2", 7, 36},
		{"Test Case 3", 11, 120},
		{"Test Case 4", 500, 76576500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Problem12(tt.input)
			if actual != tt.want {
				t.Errorf("Expected: %v, but got %v", tt.want, actual)
			}
		})
	}
}

func BenchmarkProblem12_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem12(500)
	}
}
