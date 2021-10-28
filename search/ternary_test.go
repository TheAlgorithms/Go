package search

import (
	"math"
	"testing"
)

const EPS = 1e-6

func equal(a, b float64) bool {
	return math.Abs(a-b) <= EPS
}

func TestTernaryMax(t *testing.T) {

	var tests = []struct {
		f        func(x float64) float64
		a        float64
		b        float64
		expected float64
	}{
		{f: func(x float64) float64 { return -x * x }, a: 1, b: -1, expected: 0},
		{f: func(x float64) float64 { return -2*x*x - x + 1 }, a: -1, b: 1, expected: 1.125},
	}
	for _, test := range tests {
		result, err := TernaryMax(test.a, test.b, EPS, test.f)
		if err != nil {
			t.Errorf("error occurred: %v", err)
		}
		if !equal(result, test.expected) {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
}

func TestTernaryMin(t *testing.T) {

	var tests = []struct {
		f        func(x float64) float64
		a        float64
		b        float64
		expected float64
	}{
		{f: func(x float64) float64 { return x * x }, a: -1, b: 1, expected: 0},
		{f: func(x float64) float64 { return 2*x*x + x + 1 }, a: -1, b: 1, expected: 0.875},
	}
	for _, test := range tests {
		result, err := TernaryMin(test.a, test.b, EPS, test.f)
		if err != nil {
			t.Errorf("error occurred: %v", err)
		}
		if !equal(result, test.expected) {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
}
