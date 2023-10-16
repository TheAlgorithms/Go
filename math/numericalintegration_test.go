package math

import (
	"math"
	"testing"
)

type testDataTrapezoidal struct {
	f func(float64) float64
	a float64
	b float64
	n float64
}
type testDataSimpson struct {
	f func(float64) float64
	a float64
	b float64
	n int
}

func TestTrapezoidalInteg(t *testing.T) {
	testSet := []testDataTrapezoidal{
		{func(x float64) float64 { return x * x }, 0, 1, 0.0001},
		{func(x float64) float64 { return x * x }, 10, 15, 0.0001},
		{func(x float64) float64 { return x * x * x }, 0, 1, 0.0001},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 0.0002},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 0.0001},
	}
	expectedSet := []float64{0.333, 791.666, 0.250, 0.200, 1.000}
	TOLERANCE := 0.1
	for i, test := range testSet {
		testans := TrapezoidalInteg(test.f, test.a, test.b, test.n)
		err := math.Abs(expectedSet[i] - testans)
		if err > TOLERANCE {
			t.Errorf("Tolerance=%f error=%f", TOLERANCE, err)
		}
	}
}

func BenchmarkTrapezoidalInteg(b *testing.B) {
	testSet := []testDataTrapezoidal{
		{func(x float64) float64 { return x * x }, 0, 1, 0.01},
		{func(x float64) float64 { return x * x }, 10, 15, 0.001},
		{func(x float64) float64 { return x * x * x }, 0, 1, 0.01},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 0.02},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 0.0001},
	}
	_ = []float64{0.333, 791.666, 0.250, 0.200, 1.000}
	for _, test := range testSet {
		TrapezoidalInteg(test.f, test.a, test.b, test.n)

	}
}

func TestSimpsonsOneThirdInteg(t *testing.T) {
	testSet := []testDataSimpson{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 200000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	expectedSet := []float64{0.333, 791.666, 0.250, 0.200, 1.000}
	TOLERANCE := 0.001
	for i, test := range testSet {
		testans := SimpsonsOneThirdInteg(test.f, test.a, test.b, test.n)
		err := math.Abs(expectedSet[i] - testans)
		if err > TOLERANCE {
			t.Errorf("Tolerance=%f error=%f", TOLERANCE, err)
		}
	}
}

func BenchmarkSimpsonsOneThirdInteg(b *testing.B) {
	testSet := []testDataSimpson{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 200000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	_ = []float64{0.333, 791.666, 0.250, 0.200, 1.000}

	for _, test := range testSet {
		SimpsonsOneThirdInteg(test.f, test.a, test.b, test.n)
	}
}
