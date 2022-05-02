package prime

import (
	"fmt"
	"testing"
)

var primeList = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127}
var testLimit = 127

type primalityTest func(int64) bool

func primalityTestTestingHelper(t *testing.T, name string, f primalityTest) {
	arrayIndex := 0
	for i := 1; i <= testLimit; i++ {
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

func primalityTestBenchmarkHelper(b *testing.B, f primalityTest) {
	for i := 0; i < b.N; i++ {
		f(104729)
	}
}

// Miller-Rabin Probabilistic test

func millerRabinProbabilisticTester(n int64) bool {
	result, err := MillerRabinProbabilistic(n, 40)
	if err != nil {
		panic(err)
	}

	return result
}

func TestMillerRabinProbabilistic(t *testing.T) {
	primalityTestTestingHelper(t, "Miller-Rabin Probabilistic", millerRabinProbabilisticTester)
}

func BenchmarkMillerRabinProbabilistic(b *testing.B) {
	primalityTestBenchmarkHelper(b, millerRabinProbabilisticTester)
}

// Miller-Rabin deterministic test

func millerRabinDeterministicTester(n int64) bool {
	result, err := MillerRabinDeterministic(n)
	if err != nil {
		panic(err)
	}

	return result
}

func TestMillerRabinDeterministic(t *testing.T) {
	primalityTestTestingHelper(t, "Miller-Rabin Deterministic", millerRabinDeterministicTester)
}

func BenchmarkMillerRabinDeterministic(b *testing.B) {
	primalityTestBenchmarkHelper(b, millerRabinDeterministicTester)
}

// Trial Division test

func TestTrialDivision(t *testing.T) {
	primalityTestTestingHelper(t, "Trial Division", TrialDivision)
}

func BenchmarkTrialDivision(b *testing.B) {
	primalityTestBenchmarkHelper(b, TrialDivision)
}

// Trial Division (optimized)

func TestOptimizedTrialDivision(t *testing.T) {
	primalityTestTestingHelper(t, "Trial Division (optimized)", OptimizedTrialDivision)
}

func BenchmarkOptimizedTrialDivision(b *testing.B) {
	primalityTestBenchmarkHelper(b, OptimizedTrialDivision)
}
