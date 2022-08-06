// author(s) [jo3zeph](https://github.com/jo3zeph)
// roundoff_test.go
// see roundoff.go

package roundoff

import (
	"testing"
)

func TestRoundOff(t *testing.T) {
	testCases := []struct {
		name string
		number float64
		ans float64
	}{

		//test cases 
		{
			name: "Test 1",
			number: 39.78,
			ans: 40,
		},

		{
			name: "Test 2",
			number: 12.21,
			ans: 12,
		},

		{
			name: "Test 3",
			number: -100.67,
			ans: -101,
		},
	}

	//for loop to perform all test cases
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			//pass test.number as parameter to roundoff function in roundoff.go
			returnedRoundOff := roundoff(float64(test.number))

			//log activity in console
			t.Log(test.number, " ", returnedRoundOff)

			//catch error for value mismatch between test.ans and returnedRoundOff
			if returnedRoundOff != test.ans {
				t.Errorf("Test failed. Rounded number should have been %v but received %v",
					test.ans, returnedRoundOff)
			}
		})
	}
}
	