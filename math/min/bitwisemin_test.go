// bitwisemin_test.go
// description: Test for BitwiseMin
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwisemin.go

package min

import "testing"

func TestBitwiseMin(t *testing.T) {
	base32 := 31

	t.Run("Testing(32bit) a = 32 and b = 64: ", func(t *testing.T) {
		min := BitwiseMin(32, 64, base32)
		if min != 32 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(32bit) a = 1024 and b = -9: ", func(t *testing.T) {
		min := BitwiseMin(1024, -9, base32)
		if min != -9 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(32bit) a = -6 and b = -6: ", func(t *testing.T) {
		min := BitwiseMin(-6, -6, base32)
		if min != -6 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(32bit) a = -72 and b = -73: ", func(t *testing.T) {
		min := BitwiseMin(-72, -73, base32)
		if min != -73 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	base64 := 63
	t.Run("Testing(64bit) a = 32 and b = 9223372036854775807: ", func(t *testing.T) {
		min := BitwiseMin(32, 9223372036854775807, base64)
		if min != 32 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(64bit) a = 1024 and b = -9223372036854770001: ", func(t *testing.T) {
		min := BitwiseMin(1024, -9223372036854770001, base64)
		if min != -9223372036854770001 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -6 and b = -6: ", func(t *testing.T) {
		min := BitwiseMin(-6, -6, base64)
		if min != -6 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -4223372036854775809 and b = -4223372036854775808: ", func(t *testing.T) {
		min := BitwiseMin(-4223372036854775809, -4223372036854775808, base64)
		if min != -4223372036854775809 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})

	base8 := 7
	t.Run("Testing(8bit) a = 8 and b = 16: ", func(t *testing.T) {
		min := BitwiseMin(8, 16, base8)
		if min != 8 {
			t.Fatalf("Error: BitwiseMin returned bad value")
		}
	})
}
