package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCase struct {
	price    []int
	length   int
	expected int
}

func getTestCases() []testCase {
	return []testCase{
		{[]int{0, 1, 5, 8, 9}, 4, 10},
		{[]int{0, 2, 5, 7, 8, 0}, 5, 12},
		{[]int{0, 1, 5, 8, 9, 10, 17, 17, 20}, 8, 22},
		{[]int{0, 3, 5, 8, 9, 10, 17, 17, 20}, 8, 24},
	}
}

func cutRobSolTestFunc(t *testing.T, cutRodSolFunc func([]int, int) int) {
	for _, tc := range getTestCases() {
		actual := cutRodSolFunc(tc.price, tc.length)
		if actual != tc.expected {
			t.Errorf(
				"for the price: %d and length %d expected %d, but got: %d",
				tc.price, tc.length, tc.expected, actual)
		}
	}
}

func TestCutRodRec(t *testing.T) {
	cutRobSolTestFunc(t, dynamic.CutRodRec)
}

func TestCutRodDp(t *testing.T) {
	cutRobSolTestFunc(t, dynamic.CutRodDp)
}
