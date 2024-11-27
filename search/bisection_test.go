package search

import (
	"math"
	"testing"
)

func TestBisection(t *testing.T) {
	tests := []struct {
		f       func(float64) float64
		a, b    float64
		tol     float64
		maxIter int
		want    float64
		wantErr bool
	}{
		{
			f:       func(x float64) float64 { return math.Pow(x, 3) - x - 2 },
			a:       1,
			b:       2,
			tol:     1e-6,
			maxIter: 1000,
			want:    1.521380,
			wantErr: false,
		},
		{
			f:       func(x float64) float64 { return x*x - 4 },
			a:       0,
			b:       3,
			tol:     1e-6,
			maxIter: 1000,
			want:    2,
			wantErr: false,
		},
		{
			f:       func(x float64) float64 { return x - 5 },
			a:       0,
			b:       10,
			tol:     1e-6,
			maxIter: 1000,
			want:    5,
			wantErr: false,
		},
		{
			f:       func(x float64) float64 { return math.Sin(x) },
			a:       3,
			b:       4,
			tol:     1e-6,
			maxIter: 1000,
			want:    math.Pi,
			wantErr: false,
		},
		{
			f:       func(x float64) float64 { return x*x + 1 },
			a:       -1,
			b:       1,
			tol:     1e-6,
			maxIter: 1000,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := bisection(tt.f, tt.a, tt.b, tt.tol, tt.maxIter)
		if (err != nil) != tt.wantErr {
			t.Errorf("bisection() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if math.Abs(got-tt.want) > tt.tol {
			t.Errorf("bisection() = %v, want %v", got, tt.want)
		}
	}
}
