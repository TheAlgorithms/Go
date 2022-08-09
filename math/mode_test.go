// mode.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// description: Test for Finding Mode Value In an Array

package math_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math"
)

type testCase[T constraints.Number] struct {
	name    string
	numbers []T
	mode    T
	err     error
}

func testModeFramework[T constraints.Number](t *testing.T, testCases []testCase[T]) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnMode, err := math.Mode(test.numbers)
			if returnMode != test.mode {
				t.Errorf("\n Failed test %s,\n Numbers: %v,\n Correct Mode: %v,\n Returned Mode: %v\n",
					test.name, test.numbers, test.mode, returnMode)
			}
			if !errors.Is(err, test.err) {
				t.Errorf("\n Failed test %s,\n Numbers: %v,\n Correct Error: %v,\n Returned Error: %v\n",
					test.name, test.numbers, test.err, err)
			}
		})
	}
}
func TestMode(t *testing.T) {
	// test cases for integer values
	intTestCases := []testCase[int]{
		{
			name:    "An array of positive whole numbers",
			numbers: []int{10, 52, 10, 92, 10, 75, 60, 10, 44, 29},
			mode:    10,
			err:     nil,
		},
		{
			name:    "An array of negative whole numbers",
			numbers: []int{-19, -12, -74, -19, -22, -56, -19, -19, -68, -93},
			mode:    -19,
			err:     nil,
		},
		{
			name:    "An array of positive & negative whole numbers",
			numbers: []int{18, -28, 33, -28, 2, 39, 48, -49, -87, 78, -28},
			mode:    -28,
			err:     nil,
		},
		{
			name:    "If array has no value",
			numbers: []int{},
			mode:    0,
			err:     math.ErrEmptySlice,
		},
	}
	testModeFramework(t, intTestCases)
	// test cases for float64 values
	floatTestCases := []testCase[float64]{
		{
			name:    "An array of positive real numbers",
			numbers: []float64{1.5, 2.88, 84.4, 77.2, 29.8, 46.2, 33.7, 88.4, 88.4},
			mode:    88.4,
			err:     nil,
		},
		{
			name:    "An array of negative real numbers",
			numbers: []float64{-98.1, -26.8, -54.45, -26.8, -1.5, -26.8, -33, -19.5, -26.8},
			mode:    -26.8,
			err:     nil,
		},
		{
			name:    "An array of positive and negative real numbers",
			numbers: []float64{-17, 28.9, -5.2, -19.5, 77.3, -5.2, 39.3, 28.5, -59.77, -5.2},
			mode:    -5.2,
			err:     nil,
		},
		{
			name:    "If array has no value",
			numbers: []float64{},
			mode:    0,
			err:     math.ErrEmptySlice,
		},
	}
	testModeFramework(t, floatTestCases)
}
