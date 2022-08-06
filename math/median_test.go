// author(s) [jo3zeph](https://github.com/jo3zeph)
// median_test.go
// see median.go

package math

import (
	"testing"
)

func TestMedian(t *testing.T) {
	testCases := []struct {
		name       string
		testValues []float64
		ans        float64
	}{

		//test cases
		{
			name:       "Series of numbers in ascending order",
			testValues: []float64{12, 14, 16, 18, 19},
			ans:        16,
		},

		{
			name:       "Series of numbers in random order",
			testValues: []float64{21, 10, 22, 33, 11, 88},
			ans:        21.5,
		},

		{
			name:       "Series of decimals in random order",
			testValues: []float64{11.2, 32.5, 2.5, 37.8, 21.8, 5.2},
			ans:        16.5,
		},

		{
			testValues: []float64{},
		},
	}

	//for loop to perform all test cases
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			//pass test.testValues as parameter to median function in median.go
			returnedMedian := Median(test.testValues)

			//log activity in console when error occurs
			t.Log(test.testValues, " ", returnedMedian)

			//catch error for value mismatch between test.ans and returnMedian
			if returnedMedian != test.ans {
				t.Errorf("Test failed. Median should have been %v but received %v",
					test.ans, returnedMedian)
			}
		})
	}
}
