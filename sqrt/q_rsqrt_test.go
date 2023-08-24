package sqrt

import (
	"math"
	"testing"
)

var EPSILON float32 = 0.0001

func floatEquals(a, b float32) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func Test_Q_rsqrt(t *testing.T) {
	result16 := float32(16) * Q_rsqrt(16)
	resultFromMath16 := float32(math.Sqrt(16))

	if !floatEquals(result16, resultFromMath16) {
		t.Errorf("Floating numbers %f and %f doesn't match", result16, resultFromMath16)
	}

	result4 := float32(4) * Q_rsqrt(4)
	resultFromMath4 := float32(math.Sqrt(4))

	if !floatEquals(result4, resultFromMath4) {
		t.Errorf("Floating numbers %f and %f doesn't match", result4, resultFromMath4)
	}
}
