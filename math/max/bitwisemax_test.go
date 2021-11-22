// bitwiseMax_test.go
// description: Test for Bitwise
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwiseMax.go

package max

import "testing"

func TestBitwiseMax(t *testing.T) {
	base32 := 31

	t.Run("Testing(32bit) a = 32 and m = 64: ", func(t *testing.T) {
		max := Bitwise(32, 64, base32)
		if max != 64 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(32bit) a = 1024 and m = -9: ", func(t *testing.T) {
		max := Bitwise(1024, -9, base32)
		if max != 1024 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(32bit) a = -6 and m = -6: ", func(t *testing.T) {
		max := Bitwise(-6, -6, base32)
		if max != -6 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(32bit) a = -72 and m = -73: ", func(t *testing.T) {
		max := Bitwise(-72, -73, base32)
		if max != -72 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	base64 := 63
	t.Run("Testing(64bit) a = 32 and m = 9223372036854775807: ", func(t *testing.T) {
		max := Bitwise(32, 9223372036854775807, base64)
		if max != 9223372036854775807 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = 1024 and m = -9223372036854770001: ", func(t *testing.T) {
		max := Bitwise(1024, -9223372036854770001, base64)
		if max != 1024 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -6 and m = -6: ", func(t *testing.T) {
		max := Bitwise(-6, -6, base64)
		if max != -6 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -4223372036854775809 and m = -4223372036854775808: ", func(t *testing.T) {
		max := Bitwise(-4223372036854775809, -4223372036854775808, base64)
		if max != -4223372036854775808 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	base8 := 7
	t.Run("Testing(8bit) a = 257 and m = 256: ", func(t *testing.T) {
		max := Bitwise(8, 16, base8)
		if max != 16 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})
}
