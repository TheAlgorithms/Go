package min

import (
	"testing"
)

func getTestCases() []struct {
	name    string
	base    int
	numbers []int
	min     int
} {
	var tests = []struct {
		name    string
		base    int
		numbers []int
		min     int
	}{
		{"Minimum of [128, 127], min = 117", 8, []int{128, 127}, 127},
		{"Minimum of [5], min = 5", 32, []int{5}, 5},
		{"Minimum of [-8, 32, 64, -1, 0], min = -8", 64, []int{-8, 32, 64, -1, 0}, -8},
		{"Minimum of [1, 2, 3, 4, 5], min = 1", 32, []int{1, 2, 3, 4, 5}, 1},
		{"Minimum of [1024, 512, 256, 333, 777], min = 256", 64, []int{1024, 512, 256, 333, 777}, 256},
		{"Minimum of [-9223372036854770001, -9223372036854770000, 256, 333, 777], min = 256", 64, []int{-9223372036854770001, -9223372036854770000, 256, 333, 777}, -9223372036854770001},
	}
	return tests
}

func TestBitwiseMin(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Bitwise(test.base, 999, test.numbers...)
			if result != test.min {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", test.min, result)
			}
		})
	}
}

func TestMin(t *testing.T) {
	for _, test := range getTestCases() {
		t.Run(test.name, func(t *testing.T) {
			actualMin := Int(test.numbers...)
			if actualMin != test.min {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", test.min, actualMin)
			}
		})
	}
}

func BenchmarkTestMinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(0, 10, 9, 7, 2, 1, 4, 3, 5, 6)
	}
}

func BenchmarkTestBitwiseMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bitwise(32, 10, 9, 7, 2, 1, 4, 3, 5, 6)
	}
}
