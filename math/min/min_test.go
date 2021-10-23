package min_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/min"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		left  int
		right int
		min   int
	}{
		{
			name:  "left is min",
			left:  1,
			right: 10,
			min:   1,
		},
		{
			name:  "right is min",
			left:  10,
			right: 9,
			min:   9,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMin := min.Int(test.left, test.right)
			if actualMin != test.min {
				t.Errorf("Failed test %s\n\tleft: %v, right: %v, min: %v but received: %v",
					test.name, test.left, test.right, test.min, actualMin)
			}
		})
	}
}

func TestMinOfThree(t *testing.T) {
	testCases := []struct {
		name   string
		left   int
		middle int
		right  int
		min    int
	}{
		{
			name:   "left is min",
			left:   1,
			middle: 5,
			right:  10,
			min:    1,
		},
		{
			name:   "middle is min",
			left:   10,
			middle: 5,
			right:  9,
			min:    5,
		},
		{
			name:   "right is min",
			left:   10,
			middle: 8,
			right:  6,
			min:    6,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMin := min.Int(test.left, test.middle, test.right)
			if actualMin != test.min {
				t.Errorf("Failed test %s\n\tleft: %v, middle: %v, right: %v, min: %v but received: %v",
					test.name, test.left, test.middle, test.right, test.min, actualMin)
			}
		})
	}
}
