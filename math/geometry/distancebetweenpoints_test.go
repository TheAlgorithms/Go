package geometry

import "testing"

type args struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func TestFindDistanceBetweenTwoPoints(t *testing.T) {
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"(0,0) and (2,-2)",
			args{
				0, 0, 2, -2,
			},
			2.8284271247461903,
		},
		{
			"(-20,23) and (-15,68)",
			args{
				-20, 23, -15, 68,
			},
			45.27692569068709,
		},
		{
			"(2,2) and (14,11)",
			args{
				2, 2, 14, 11,
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDistanceBetweenTwoPoints(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("FindDistanceBetweenTwoPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindDistanceBetweenTwoPoints(b *testing.B) {
	benchmarks := []struct {
		name     string
		args     args
		expected float64
	}{
		{
			"(0,0) and (2,-2)",
			args{
				0, 0, 2, -2,
			},
			2.8284271247461903,
		},
		{
			"(-20,23) and (-15,68)",
			args{
				-20, 23, -15, 68,
			},
			45.27692569068709,
		},
		{
			"(2,2) and (14,11)",
			args{
				2, 2, 14, 11,
			},
			15,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindDistanceBetweenTwoPoints(bm.args.x1, bm.args.y1, bm.args.x2, bm.args.x2)
			}
		})
	}
}
