package math

import (
	"math"
	"testing"
)

type testData struct{
	f func(float64) float64
	a float64
	b float64
	n int
}


func TestTrapezoidalInteg(t *testing.T){
	testSet := []testData{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 20000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	expectedSet := []float64{0.333,791.666, 0.250, 0.200, 1.000}
	TOLERANCE := 0.1
	for i,test:=range testSet{
		testans:=TrapezoidalInteg(test.f,test.a,test.b,test.n)
		err:=math.Abs(expectedSet[i]-testans)
		if err > TOLERANCE{
			t.Errorf("Tolerance=%f error=%f",TOLERANCE,err)
		}
	}
}



func BenchmarkTrapezoidalInteg(b *testing.B){
	testSet := []testData{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 20000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	_= []float64{0.333,791.666, 0.250, 0.200, 1.000}
	for _,test:=range testSet{
		TrapezoidalInteg(test.f,test.a,test.b,test.n)
		
	}
}

func TestSimpsonsOneThirdInteg(t *testing.T){
	testSet := []testData{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 200000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	expectedSet := []float64{0.333,791.666, 0.250, 0.200, 1.000}
	TOLERANCE := 0.001
	for i,test:=range testSet{
		testans:=SimpsonsOneThirdInteg(test.f,test.a,test.b,test.n)
		err:=math.Abs(expectedSet[i]-testans)
		if err > TOLERANCE{
			t.Errorf("Tolerance=%f error=%f",TOLERANCE,err)
		}
	}
}

func BenchmarkSimpsonsOneThirdInteg(b *testing.B) {
	testSet := []testData{
		{func(x float64) float64 { return x * x }, 0, 1, 200},
		{func(x float64) float64 { return x * x }, 10, 15, 200000},
		{func(x float64) float64 { return x * x * x }, 0, 1, 600},
		{func(x float64) float64 { return x * x * x * x }, 0, 1, 1100},
		{func(x float64) float64 { return math.Sin(x) }, 0, math.Pi / 2, 2000},
	}
	_ = []float64{0.333,791.666, 0.250, 0.200, 1.000}
	
	for _,test:=range testSet{
		SimpsonsOneThirdInteg(test.f,test.a,test.b,test.n)
	}
}