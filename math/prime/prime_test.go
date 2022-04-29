package prime

import (
	"fmt"
	"testing"
)

var primeList = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
var testLimit = 127

func PrimalityTestTestingHelper(t *testing.T, name string, f func(int64 n) bool) {
	arrayIndex := 0
	for i := 1; i < testLimit; i++ {
		isPrime := i == primeList[arrayIndex]

		testName := fmt.Sprintf("%s(%d)", name, i)
		t.Run(testName, func(t *testing.T) {
			result := f(int64(i))

			if isPrime {
				arrayIndex++
			}

			if result != isPrime {
				t.Errorf("%d: %s function returned %v\n", i, name, result)
			}
		})
	}
}
