package math

import (
	"testing"
)

type SumOfProperDivisorsExample struct {
	input    uint
	expected uint
}

// getSumOfProperDivisorsExamples returns an array of example data
// for the tests of the function SumOfProperDivisors.
// Example data was verified using [A001065].
// [A001065]: https://oeis.org/A001065/b001065.txt
func getSumOfProperDivisorsExamples() []SumOfProperDivisorsExample {
	return []SumOfProperDivisorsExample{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 1},
		{6, 6},
		{7, 1},
		{8, 7},
		{9, 4},
		{10, 8},
		{11, 1},
		{12, 16},
		{13, 1},
		{14, 10},
		{15, 9},
		{16, 15},
		{17, 1},
		{18, 21},
		{19, 1},
		{20, 22},
		{21, 11},
		{22, 14},
		{23, 1},
		{24, 36},
		{25, 6},
		{26, 16},
		{27, 13},
		{28, 28},
		{29, 1},
		{30, 42},
		{31, 1},
		{32, 31},
		{33, 15},
		{34, 20},
		{35, 13},
		{36, 55},
		{37, 1},
		{38, 22},
		{39, 17},
		{40, 50},
		{41, 1},
		{42, 54},
		{43, 1},
		{44, 40},
		{45, 33},
		{46, 26},
		{47, 1},
		{48, 76},
		{49, 8},
		{50, 43},
		{51, 21},
		{52, 46},
		{53, 1},
		{54, 66},
		{55, 17},
		{56, 64},
		{57, 23},
		{58, 32},
		{59, 1},
		{60, 108},
		{800, 1153},
		{8204, 8260},
		{9646, 8498},
		{9756, 14996},
		{9829, 1},
		{9954, 15006},
		{10000, 14211},
	}
}

func TestSumOfProperDivisors(t *testing.T) {
	for _, tc := range getSumOfProperDivisorsExamples() {
		actual := SumOfProperDivisors(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected SumOfProperDivisors(%d) to be: %d, but got: %d", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkSumOfProperDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfProperDivisors(800)
	}
}

// getPerfectNumberExamples returns an array of some perfect numbers.
// Example data was verified using [A000396].
// [A000396]: https://oeis.org/A000396
func getPerfectNumberExamples() []uint {
	return []uint{
		6,
		28,
		496,
		8128,
		33550336}
}

func TestIsPerfectNumber(t *testing.T) {
	for _, tc := range getPerfectNumberExamples() {
		if !IsPerfectNumber(tc) {
			t.Errorf("%d is expected to be a perfect number.", tc)
		}
	}
}

func getNotPerfectNumberExamples() []uint {
	return []uint{
		0,
		1,
		2,
		3,
		4,
		5,
		7,
		100,
		219,
		997,
		33550335,
		33550337,
	}
}

func TestIsNotPerfectNumber(t *testing.T) {
	for _, tc := range getNotPerfectNumberExamples() {
		if IsPerfectNumber(tc) {
			t.Errorf("%d is expected not to be a perfect number.", tc)
		}
	}
}

func BenchmarkIsPerfectNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPerfectNumber(8128)
	}
}
