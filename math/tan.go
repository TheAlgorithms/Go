package math

import (
	"math"
)

func Tan(x float64) float64 {
	return Cot((math.Pi / 2) - x)
}
