package binarymath

import (
	"math"
	"testing"
)

func getTestsForPowerOfTwo() []struct {
	name    string
	a       int
	missing bool
} {
	var tests = []struct {
		name    string
		a       int
		missing bool
	}{
		{"64 is power of two? - YES", 64, true},
		{"1 is power of two? - YES", 1, true},
		{"2 is power of two? - YES", 2, true},
		{"5 is power of two? - NO", 5, false},
		{"1023 is power of two? - NO", 1023, false},
		{"1024 is power of two? - YES", 1024, true},
		{"0 is power of two? - NO", 0, false},
		{"9223372036854775807 is power of two? - NO", math.MaxInt64, false},
		{"9223372036854775806 is power of two? - NO", math.MaxInt64, false},
		{"4611686018427387904 is power of two? - YES", 4611686018427387904, true},
	}
	return tests
}

func TestIsPowerOfTwoBinaryMethod(t *testing.T) {
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IsPowerOfTwoBinaryMethod(tv.a)
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
}

func TestIsPowerOfTwoUseCycleAndLeftShift(t *testing.T) {
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IsPowerOfTwoUseCycleAndLeftShift(uint(tv.a))
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
}

func TestIsPowerOfTwoUseLog(t *testing.T) {
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IsPowerOfTwoUseLog(float64(tv.a))
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
}

func BenchmarkIsPowerOfTwoBinaryMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPowerOfTwoBinaryMethod(1024)
	}
}

func BenchmarkIsPowerOfTwoUseCycleAndLeftShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPowerOfTwoUseCycleAndLeftShift(1024)
	}
}

func BenchmarkIsPowerOfTwoUseLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPowerOfTwoUseLog(1024)
	}
}
