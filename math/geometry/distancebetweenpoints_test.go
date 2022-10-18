package geometry

import (
	"testing"
)

type args struct {
	p1 CartesianPoint
	p2 CartesianPoint
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
				CartesianPoint{0, 0},
				CartesianPoint{2, -2},
			},
			2.8284271247461903,
			false,
		},
		{
			"(-20,23) and (-15,68)",
			args{
				CartesianPoint{-20, 23},
				CartesianPoint{-15, 68},
			},
			45.27692569068709,
			false,
		},
		{
			"(2,2) and (14,11)",
			args{
				CartesianPoint{2, 2},
				CartesianPoint{14, 11},
			},
			15,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDistanceBetweenTwoPoints(tt.args.p1, tt.args.p2)
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
	benchmarks := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			"(0,0) and (2,-2)",
			args{
				CartesianPoint{0, 0},
				CartesianPoint{2, -2},
			},
			2.8284271247461903,
			false,
		},
		{
			"(-20,23) and (-15,68)",
			args{
				CartesianPoint{-20, 23},
				CartesianPoint{-15, 68},
			},
			45.27692569068709,
			false,
		},
		{
			"(2,2) and (14,11)",
			args{
				CartesianPoint{2, 2},
				CartesianPoint{14, 11},
			},
			15,
			false,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindDistanceBetweenTwoPoints(bm.args.p1, bm.args.p2)
			}
		})
	}
}
