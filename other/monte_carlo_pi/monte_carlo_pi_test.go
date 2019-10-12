package main

import (
	"math"
	"testing"
)

func TestMonteCarloPi(t *testing.T) {
	delta := 0.0001

	pi := MonteCarloPi(100000000)

	if math.Abs(pi-math.Pi) > delta {
		t.Errorf("Given: %.4f, expected: %.4f", pi, math.Pi)
	}
}
