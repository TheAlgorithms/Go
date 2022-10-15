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
		{"314", 3},
		{"31415", 5},
		{"314159", 6},
		{"31415926535", 11},
		{"314159265358", 12},
		{"314159265358979323846", 21},
		{"3141592653589793238462643", 25},
		{"314159265358979323846264338327950", 33},
		{"31415926535897932384626433832795028841971693993751" +
			"05820974944592307816406286208998628034825342117067" +
			"98214808651328230664709384460955058223172535940812" +
			"84811174502841027019385211055596446229489549303819" +
			"64428810975665933446128475648233786783165271201909" +
			"14564856692346034861045432664821339360726024914127" +
			"37245870066063155881748815209209628292540917153643" +
			"678925903600", 362},
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
