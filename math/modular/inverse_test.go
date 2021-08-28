// inverse_test.go
// description: Test for Modular Inverse
// author(s) [Taj](https://github.com/tjgurwara99)
// see inverse.go

package modular

import "testing"

func TestInverse(t *testing.T) {
	t.Run("Testing a = 3 and m = 11: ", func(t *testing.T) {
		inv, err := Inverse(3, 11)
		if err != nil {
			t.Fatalf("Error was raised when it shouldn't: %v", err)
		}
		if inv != 4 {
			t.Fatalf("Test failed:\n\tExpected values: %v\n\tReturned value: %v", 4, inv)
		}
	})
	t.Run("Testing a = 10 and m = 17: ", func(t *testing.T) {
		inv, err := Inverse(10, 17)
		if err != nil {
			t.Fatalf("Error was raised when it shouldn't: %v", err)
		}
		if inv != 12 {
			t.Fatalf("Test failed:\n\tExpected values: %v\n\tReturned value: %v", 12, inv)
		}
	})
}
