package dynamic

import (
	"fmt"
	"testing"
)

func TestCatalanNumbers(t *testing.T) {
	var testCatalanNumbersData = []struct {
		nthCatalanNumber      int
		expectedCatalanNumber int64
		expectedError         error
	}{
		{nthCatalanNumber: -1000, expectedCatalanNumber: 0, expectedError: errCatalan},
		{nthCatalanNumber: -1, expectedCatalanNumber: 0, expectedError: errCatalan},
		{nthCatalanNumber: 0, expectedCatalanNumber: 1, expectedError: nil},
		{nthCatalanNumber: 1, expectedCatalanNumber: 1, expectedError: nil},
		{nthCatalanNumber: 2, expectedCatalanNumber: 2, expectedError: nil},
		{nthCatalanNumber: 3, expectedCatalanNumber: 5, expectedError: nil},
		{nthCatalanNumber: 4, expectedCatalanNumber: 14, expectedError: nil},
		{nthCatalanNumber: 5, expectedCatalanNumber: 42, expectedError: nil},
		{nthCatalanNumber: 6, expectedCatalanNumber: 132, expectedError: nil},
		{nthCatalanNumber: 7, expectedCatalanNumber: 429, expectedError: nil},
		{nthCatalanNumber: 8, expectedCatalanNumber: 1430, expectedError: nil},
		{nthCatalanNumber: 9, expectedCatalanNumber: 4862, expectedError: nil},
		{nthCatalanNumber: 10, expectedCatalanNumber: 16796, expectedError: nil},
		{nthCatalanNumber: 1000, expectedCatalanNumber: 4233371109654655040, expectedError: nil},
	}

	for i := range testCatalanNumbersData {
		t.Run(fmt.Sprintf("the %dth Catalan Number", testCatalanNumbersData[i].nthCatalanNumber), func(t *testing.T) {
			nthCatalanNumber := testCatalanNumbersData[i].nthCatalanNumber
			result, err := NthCatalanNumber(nthCatalanNumber)
			expectedCatalanNumber := testCatalanNumbersData[i].expectedCatalanNumber
			expectedError := testCatalanNumbersData[i].expectedError

			if err != expectedError {
				t.Errorf("Expected %dth Catalan Number error: %d\nFound: %d\n", nthCatalanNumber, expectedError, err)
			}

			if result != expectedCatalanNumber {
				t.Errorf("Expected %dth Catalan Number: %d\nFound: %d\n", nthCatalanNumber, expectedCatalanNumber, result)
			}
		})
	}
}
