// extended_test.go
// description: Test for Extended GCD algorithm in extended.go
// author(s) [Taj](https://github.com/tjgurwara99)
// see extended.go

package gcd

import "testing"

func TestExtended(t *testing.T) {
	var testCasesExtended = []struct {
		name string
		a    int64
		b    int64
		gcd  int64
		x    int64
		y    int64
	}{
		{"gcd of 30 and 50", 30, 50, 10, 2, -1},
	}
	for _, tc := range testCasesExtended {
		t.Run(tc.name, func(t *testing.T) {
			gcd, x, y := Extended(tc.a, tc.b)
			if gcd != tc.gcd && x != tc.x && y != tc.y {
				t.Fatalf("Expected values:\n\tGCD: Expected %v Returned %v,\n\tx: Expected %v Returned %v\n\ty: Expected %v Returned %v", tc.gcd, gcd, tc.x, x, tc.y, y)
			}
		})
	}
}
