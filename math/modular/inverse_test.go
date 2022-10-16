// inverse_test.go
// description: Test for Modular Inverse
// author(s) [Taj](https://github.com/tjgurwara99)
// see inverse.go

package modular

import "testing"
import "fmt"

func TestInverse(t *testing.T) {
	testCases := []struct {
		a             int64
		m             int64
		expectedValue int64
		expectedError error
	}{
		{3, 11, 4, nil},
		{10, 17, 12, nil},
		{2, 6, 0, ErrorInverse},
		{1, 0, 0, ErrorInverse},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("Testing a = %d and m = %d: ", tc.a, tc.m)
		t.Run(testName, func(t *testing.T) {
			inv, err := Inverse(tc.a, tc.m)
			if err != tc.expectedError {
				if tc.expectedError == nil {
					t.Fatalf("Error was raised when it shouldn't: %v", err)
				} else {
					t.Fatalf("Error was not raised when it should")
				}
			}
			if inv != tc.expectedValue {
				t.Fatalf("expected: %d, got: %d", tc.expectedValue, inv)
			}
		})
	}
}
