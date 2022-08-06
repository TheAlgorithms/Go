// mode.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// description: Test for Finding Mode Value In an Array

package mode

import "testing"

func TestMode(t *testing.T) {

	testCases := []struct {
		name    string
		numbers []float64
		mode    float64
	}{
		// Test Cases
		{
			name:    "An array of positive whole numbers",
			numbers: []float64{10, 52, 10, 92, 10, 75, 60, 10, 44, 29},
			mode:    10,
		},
		{
			name:    "An array of negative whole numbers",
			numbers: []float64{-19, -12, -74, -19, -22, -56, -19, -19, -68, -93},
			mode:    -19,
		},
		{
			name:    "An array of positive & negative whole numbers",
			numbers: []float64{18, -28, 33, -28, 2, 39, 48, -49, -87, 78, -28},
			mode:    -28,
		},
		{
			name:    "An array of positive real numbers",
			numbers: []float64{1.5, 2.88, 84.4, 77.2, 29.8, 46.2, 33.7, 88.4, 88.4},
			mode:    88.4,
		},
		{
			name:    "An array of negative real numbers",
			numbers: []float64{-98.1, -26.8, -54.45, -26.8, -1.5, -26.8, -33, -19.5, -26.8},
			mode:    -26.8,
		},
		{
			name:    "An array of positve and negative real numbers",
			numbers: []float64{-17, 28.9, -5.2, -19.5, 77.3, -5.2, 39.3, 28.5, -59.77, -5.2},
			mode:    -5.2,
		},
	}

	// Run for all the test cases
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnMode := Mode(test.numbers)
			if returnMode != test.mode {
				t.Errorf("\n Failed test %s,\n Numbers: %v,\n Correct Mode: %v,\n Returned Mode: %v\n",
					test.name, test.numbers, test.mode, returnMode)
			}
		})
	}
}
