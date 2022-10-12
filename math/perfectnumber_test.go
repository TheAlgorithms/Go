package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

type testCaseSumOfProperDivisors struct {
	input    uint
	expected uint
}

// getSumOfProperDivisorsTestCases returns an array of test data
// for the tests of the function SumOfProperDivisors.
// Data was verified using [A001065].
// [A001065]: https://oeis.org/A001065/b001065.txt
func getSumOfProperDivisorsTestCases() []testCaseSumOfProperDivisors {
	return []testCaseSumOfProperDivisors{
		{0, 0},
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
	for _, tc := range getSumOfProperDivisorsTestCases() {
		actual := math.SumOfProperDivisors(tc.input)
		if actual != tc.expected {
			t.Errorf(
				"Expected SumOfProperDivisors(%d) to be: %d, but got: %d",
				tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkSumOfProperDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.SumOfProperDivisors(800)
	}
}

type isPerfectNumberTestCase struct {
	number    uint
	isPerfect bool
}

// getIsPerfectNumberTestCases returns an array of test data
// for the tests of the function IsPerfectNumber
// Data was verified using [A000396].
// [A000396]: https://oeis.org/A000396
func getIsPerfectNumberTestCases() []isPerfectNumberTestCase {
	return []isPerfectNumberTestCase{
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{33550336, true},
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{7, false},
		{100, false},
		{219, false},
		{997, false},
		{33550335, false},
		{33550337, false},
	}
}

func TestIsPerfectNumber(t *testing.T) {
	for _, tc := range getIsPerfectNumberTestCases() {
		if math.IsPerfectNumber(tc.number) != tc.isPerfect {
			t.Errorf("Unexpected output for %d", tc.number)
		}
	}
}

func BenchmarkIsPerfectNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.IsPerfectNumber(8128)
	}
}
