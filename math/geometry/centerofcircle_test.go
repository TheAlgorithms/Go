// centerofcircle_test.go
// description: Center of the circle
// author(s) [red_byte](https://github.com/i-redbyte)
// see centerofcircle.go

// Package geometry contains the solution of geometric problems
package geometry

import (
	"fmt"
	"testing"
)

func TestMonteCarloPi(t *testing.T) {
	t.Run("Find the center of the circle using endpoints of diameter", func(t *testing.T) {
		x, y := FindCenterUsEndpointsDiameter(-11, 5, 6, -8)
		result := fmt.Sprintf("(%.2f:%.2f)", x, y)
		t.Log(result)
		if result != "(-2.50:-1.50)" {
			t.Errorf("Wrong result! Expected:%f, returned:%s ", 3.1415, result)
		}
	})
}

func BenchmarkMonteCarloPi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindCenterUsEndpointsDiameter(-9, 3, 5, -7)
	}
}
