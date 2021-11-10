package max

import "testing"

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		left  int
		right int
		max   int
	}{
		{
			name:  "Left is max",
			left:  10,
			right: 9,
			max:   10,
		},
		{
			name:  "right is max",
			left:  1,
			right: 10,
			max:   10,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnedMax := Int(test.left, test.right)
			if returnedMax != test.max {
				t.Errorf("Failed test %s\n\tleft: %v, right: %v, max: %v but received: %v",
					test.name, test.left, test.right, test.max, returnedMax)
			}
		})
	}
}

func TestMaxOfThree(t *testing.T) {
	testCases := []struct {
		name   string
		left   int
		middle int
		right  int
		max    int
	}{
		{
			name:   "right is max",
			left:   1,
			middle: 5,
			right:  10,
			max:    10,
		},
		{
			name:   "left is max",
			left:   10,
			middle: 5,
			right:  9,
			max:    10,
		},
		{
			name:   "left is max",
			left:   10,
			middle: 8,
			right:  6,
			max:    10,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMax := Int(test.left, test.middle, test.right)
			if actualMax != test.max {
				t.Errorf("Failed test %s\n\tleft: %v, middle: %v, right: %v, max: %v but received: %v",
					test.name, test.left, test.middle, test.right, test.max, actualMax)
			}
		})
	}
}
