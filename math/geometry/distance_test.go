package geometry

import (
	"testing"
)

type args struct {
	p1 EuclideanPoint
	p2 EuclideanPoint
}

func TestFindDistanceBetweenTwoPoints(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			"(0,0) and (2,-2)",
			args{
				EuclideanPoint{0, 0},
				EuclideanPoint{2, -2},
			},
			2.8284271247461903,
			false,
		},
		{
			"(-20,23) and (-15,68)",
			args{
				EuclideanPoint{-20, 23},
				EuclideanPoint{-15, 68},
			},
			45.27692569068709,
			false,
		},
		{
			"(2,2) and (14,11)",
			args{
				EuclideanPoint{2, 2},
				EuclideanPoint{14, 11},
			},
			15,
			false,
		},
		{
			"Return error for mismatched dimensions(2,2) and ()",
			args{
				EuclideanPoint{2, 2},
				EuclideanPoint{},
			},
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EuclideanDistance(tt.args.p1, tt.args.p2)
			if (err != nil) != tt.wantErr {
				t.Errorf("EuclideanDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EuclideanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindDistanceBetweenTwoPoints(b *testing.B) {
	p1 := EuclideanPoint{0, 0}
	p2 := EuclideanPoint{2, -2}

	for i := 0; i < b.N; i++ {
		_, _ = EuclideanDistance(p1, p2)
	}
}
