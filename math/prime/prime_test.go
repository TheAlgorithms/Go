package prime

import (
	"fmt"
	"testing"
)

var primeList = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
var testLimit = 127

type primalityTest func(int64 n) (bool, error)

func primalityTestTestingHelper(t *testing.T, name string, f primalityTest) {
	arrayIndex := 0
	for i := 1; i < testLimit; i++ {
		isPrime := i == primeList[arrayIndex]

		testName := fmt.Sprintf("%s(%d)", name, i)
		t.Run(testName, func(t *testing.T) {
			result, err := f(int64(i))
			if err != nil {
				t.Error(err)
			}

			if isPrime {
				arrayIndex++
			}

			if result != isPrime {
				t.Errorf("%d: %s function returned %v\n", i, name, result)
			}
		})
	}
}

func primalityTestBenchmarkHelper(b *testing.B, f primalityTest) {
	for i := 0; i < b.N; i++ {
		f(104729)
	}
}

// Miller-Rabin Probabilistic test

func millerRabinProbabilisticTester(n int64) (bool, error) {
	return MillerRabinProbabilistic(n, 40)
}

func TestMillerRabinProbabilistic(t *testing.T) {
	primalityTestTestingHelper(t, millerRabinProbailisticTester)
}

func BenchmarkMillerRabinProbabilistic(b *testing.B) {
	primalityTestBenchmarkHelper(b, millerRabinProbabilisticTester)
}

// Miller-Rabin deterministic test

func TestMillerRabinDeterministic(t *testing.T) {
	primalityTestTestingHelper(t, MillerRabinDeterministic)
}

func BenchmarkMillerRabinDeterministic(b *testing.B) {
	primalityTestBenchmarkHelper(b, MillerRabinDeterministic)
}
