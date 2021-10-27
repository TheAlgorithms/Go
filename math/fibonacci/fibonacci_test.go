package fibonacci

import (
	"testing"
)

func TestNthfib(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		result int
	}{
		{
			name:   "1st Fibonacci number",
			n:      1,
			result: 1,
		},
		{
			name:   "3rd Fibonacci number",
			n:      3,
			result: 2,
		},
		{
			name:   "10th Fibonacci number",
			n:      10,
			result: 55,
		},
		{
			name:   "Invalid",
			n:      0,
			result: 0,
		},
	}

	for _, tc := range testCases {
		actual_result := Nthfib(tc.n)
		if actual_result != tc.result {
			t.Errorf("Expected %dth Fibonacci is %d, but got %d", tc.n, tc.result,
				actual_result)
		}
	}

}
