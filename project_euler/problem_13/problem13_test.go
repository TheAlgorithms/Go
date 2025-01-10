package problem13

import "testing"

func TestProblem13_Func(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Test Case 1", "5537376230"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Problem13()
			if actual != tt.expected {
				t.Errorf("Expected: %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func BenchmarkProblem13_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem13()
	}
}
