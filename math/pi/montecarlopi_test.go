// montecarlopi_test.go
// description: Test for calculating pi by the Monte Carlo method
// author(s) [red_byte](https://github.com/i-redbyte)
// see montecarlopi.go

package pi

import (
	"fmt"
	"testing"
)

func TestMonteCarloPi(t *testing.T) {
	t.Run("Monte Carlo Pi", func(t *testing.T) {
		result := fmt.Sprintf("%.2f", MonteCarloPi(100000000))
		t.Log(result)
		if result != "3.14" {
			t.Errorf("Wrong result! Expected:%f, returned:%s ", 3.1415, result)
		}
	})
}

func BenchmarkMonteCarloPi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonteCarloPi(100000000)
	}
}
