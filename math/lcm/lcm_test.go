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
		// Regression: integer-arithmetic fix for float precision loss and overflow.
		{
			// Mode A: |a*b| overflows int64 even though the true LCM fits.
			// Previously returned 0 (impossible for positive inputs).
			name:   "LCM of 2^32 & 2^33 (overflow-safe intermediate)",
			a:      4294967296,
			b:      8589934592,
			output: 8589934592,
		}, {
			// Mode A: lcm(x, x) == x. Previously lcm(x,x) != x for large x.
			name:   "LCM of 3037000500 & 3037000500 (lcm(x,x)==x)",
			a:      3037000500,
			b:      3037000500,
			output: 3037000500,
		}, {
			// Mode B: float64 cannot represent 2^53+1 exactly.
			// Previously returned 9007199254740992, violating lcm(x,1)==x.
			name:   "LCM of 9007199254740993 & 1 (2^53+1 precision)",
			a:      9007199254740993,
			b:      1,
			output: 9007199254740993,
		}, {
			// Mode B: product exceeds 2^53 but fits int64; float64 loses low bits.
			// Previously returned 1000000016000000000 (off by 63).
			name:   "LCM of 1000000007 & 1000000009 (large coprime precision)",
			a:      1000000007,
			b:      1000000009,
			output: 1000000016000000063,
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
