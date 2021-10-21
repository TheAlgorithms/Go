package dynamic

import (
	"fmt"
	"testing"
)

func TestCatalanNumbers(t *testing.T) {
	var testCatalanNumbersData = []struct {
		nthCatalanNumber      int
		expectedCatalanNumber int64
	}{
		{nthCatalanNumber: -1000, expectedCatalanNumber: -1},
		{nthCatalanNumber: -1, expectedCatalanNumber: -1},
		{nthCatalanNumber: 0, expectedCatalanNumber: 1},
		{nthCatalanNumber: 1, expectedCatalanNumber: 1},
		{nthCatalanNumber: 2, expectedCatalanNumber: 2},
		{nthCatalanNumber: 3, expectedCatalanNumber: 5},
		{nthCatalanNumber: 4, expectedCatalanNumber: 14},
		{nthCatalanNumber: 5, expectedCatalanNumber: 42},
		{nthCatalanNumber: 6, expectedCatalanNumber: 132},
		{nthCatalanNumber: 7, expectedCatalanNumber: 429},
		{nthCatalanNumber: 8, expectedCatalanNumber: 1430},
		{nthCatalanNumber: 9, expectedCatalanNumber: 4862},
		{nthCatalanNumber: 10, expectedCatalanNumber: 16796},
		{nthCatalanNumber: 1000, expectedCatalanNumber: 4233371109654655040},
	}

	for i := range testCatalanNumbersData {
		t.Run(fmt.Sprintf("the %dth Catalan Number", testCatalanNumbersData[i].nthCatalanNumber), func(t *testing.T) {
			nthCatalanNumber := testCatalanNumbersData[i].nthCatalanNumber
			result := NthCatalanNumber(nthCatalanNumber)
			expectedCatalanNumber := testCatalanNumbersData[i].expectedCatalanNumber

			if result != expectedCatalanNumber {
				t.Errorf("Expected %dth Catalan Number: %d\nFound: %d\n", nthCatalanNumber, expectedCatalanNumber, result)
			}
		})
	}
}
