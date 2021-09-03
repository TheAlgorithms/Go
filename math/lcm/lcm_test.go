package lcm

import "testing"

func TestLcm(t *testing.T) {
	testCases := []struct {
		name   string
		a      int64
		b      int64
		output int64
	}{
		{
			name:   "LCM of 1 & 5",
			a:      1,
			b:      5,
			output: 5,
		}, {
			name:   "LCM of 2 & 5",
			a:      2,
			b:      5,
			output: 10,
		}, {
			name:   "LCM of 5 & 10",
			a:      10,
			b:      5,
			output: 10,
		}, {
			name:   "LCM of 5 & 5",
			a:      5,
			b:      5,
			output: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual_output := Lcm(tc.a, tc.b)
			if actual_output != tc.output {
				t.Errorf("Expected LCM of %d and %d is %d, but got %d", tc.a, tc.b, tc.output,
					actual_output)
			}
		})
	}
}
