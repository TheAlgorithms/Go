package dynamic

import "testing"

func TestSubsetSum(t *testing.T) {

	var subsetSumTestData = []struct {
		description    string
		array          []int
		sum            int
		expectedResult bool
		expectedError  error
	}{
		{
			description:    "array of size 0 and sum 0",
			array:          []int{},
			sum:            0,
			expectedResult: true,
			expectedError:  nil,
		},
		{
			description:    "array of size 0 and non-zero sum",
			array:          []int{},
			sum:            2,
			expectedResult: false,
			expectedError:  nil,
		},
		{
			description:    "array of size 2 and sum 0",
			array:          []int{1, 2},
			sum:            0,
			expectedResult: true,
			expectedError:  nil,
		},
		{
			description:    "array of size 5 and sum 6",
			array:          []int{1, 2, 3, 4, 5},
			sum:            6,
			expectedResult: true,
			expectedError:  nil,
		},
		{
			description:    "array of size 5 and sum 7",
			array:          []int{1, 2, 3, 4, 5},
			sum:            7,
			expectedResult: true,
			expectedError:  nil,
		},
		{
			description:    "array of size 5 and sum 17",
			array:          []int{1, 2, 3, 4, 5},
			sum:            17,
			expectedResult: false,
			expectedError:  nil,
		},
		{
			description:    "array of size 6 negative values and sum positive",
			array:          []int{-1, -2, -3, -4, -5, -6},
			sum:            1,
			expectedResult: false,
			expectedError:  ErrInvalidPosition,
		},
		{
			description:    "array of size 6 with positive and negative values, also with a negative sum",
			array:          []int{-1, -1, 3, 4, 5, 6},
			sum:            -2,
			expectedResult: false,
			expectedError:  ErrNegativeSum,
		},
	}

	for i := range subsetSumTestData {

		t.Run(subsetSumTestData[i].description, func(t *testing.T) {
			array := subsetSumTestData[i].array
			sum := subsetSumTestData[i].sum

			expectedResult := subsetSumTestData[i].expectedResult
			result, err := IsSubsetSum(array, sum)
			expectedError := subsetSumTestData[i].expectedError

			if err != expectedError {
				t.Logf("FAIL: %s", subsetSumTestData[i].description)
				t.Fatalf("Expected error:%t\nFound: %t", expectedError, err)
			}

			if result != expectedResult {
				t.Logf("FAIL: %s", subsetSumTestData[i].description)
				t.Fatalf("Expected result:%t\nFound: %t", expectedResult, result)
			}
		})
	}
}
