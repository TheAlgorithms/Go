// spigotpi_test.go
// description: Test for Spigot Algorithm for the Digits of Pi
// author(s) [red_byte](https://github.com/i-redbyte)
// see spigotpi.go
package pi

import (
	"strconv"
	"testing"
)

func TestSpigot(t *testing.T) {
	var tests = []struct {
		result string
		n      int
	}{
		{"314159", 6},
		{"31415926535", 11},
		{"314159265358", 12},
		{"314159265358979323846", 21},
		{"314", 3},
		{"31415", 5},
	}
	for _, tv := range tests {
		t.Run(strconv.Itoa(tv.n)+":"+tv.result, func(t *testing.T) {
			result := Spigot(tv.n)
			if result != tv.result {
				t.Errorf("Bad result %d:%s", tv.n, tv.result)
			}
		})
	}
}

func BenchmarkPiSpigotN10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Spigot(10)
	}
}

func BenchmarkPiSpigotN100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Spigot(100)
	}
}

func BenchmarkPiSpigotN1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Spigot(1000)
	}
}
