package problem18

import "testing"

// Tests
func TestProblem18_Func(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		deep     int
		expected int
	}{
		{
			name:     "Test case 1",
			input:    problem18_test_parsed_string,
			deep:     2,
			expected: 23,
		},
		{
			name:     "Test case 2",
			input:    problem18_input_parsed_string,
			deep:     2,
			expected: 1074,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Problem18(test.input, test.deep)
			if actual != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, actual)
			}
		})
	}
}

// Benchmarks
func BenchmarkProblem18_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem18(problem18_input_parsed_string, 2)
	}
}
