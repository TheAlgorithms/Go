package dynamic

import (
	"testing"
)

var rodCuttingExamples = []struct {
	price    []int
	length   int
	expected int
}{
	{[]int{0, 1, 5, 8, 9}, 4, 10},
	{[]int{0, 2, 5, 7, 8, 0}, 5, 12},
	{[]int{0, 1, 5, 8, 9, 10, 17, 17, 20}, 8, 22},
	{[]int{0, 3, 5, 8, 9, 10, 17, 17, 20}, 8, 24},
}

func CutRobSolutionFunctionTemplate(t *testing.T, cutRodSolutionFunction func([]int, int) int) {
	for _, tc := range rodCuttingExamples {
		actual := cutRodSolutionFunction(tc.price, tc.length)
		if actual != tc.expected {
			t.Errorf("for the price: %d and length %d expected %d, but got: %d", tc.price, tc.length, tc.expected, actual)
		}
	}
}

func TestCutRodRec(t *testing.T) {
	CutRobSolutionFunctionTemplate(t, CutRodRec)
}

func TestCutRodDp(t *testing.T) {
	CutRobSolutionFunctionTemplate(t, CutRodDp)
}
