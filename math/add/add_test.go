package add

import "testing"

func TestAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		a, b     int
		expected int
	}{
		{5, 3, 8},
		{-2, 7, 5},
		{0, 0, 0},
		{10, -5, 5},
<<<<<<< HEAD
=======
		{1, 1, 2},
>>>>>>> 28a14418d73e15b9b58a21cd08ca31884fdee9a8
	}

	// Iterate over the test cases
	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("For %d + %d, expected %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}
