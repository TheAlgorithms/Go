// montecarlopi_test.go
// description: Test for calculating pi by the Monte Carlo method
// author(s) [red_byte](https://github.com/i-redbyte)
// see montecarlopi.go

package pi

import (
	"fmt"
	"testing"
)

func TestMonteCarloPi(t *testing.T) {
	t.Run("Monte Carlo Pi", func(t *testing.T) {
		result := fmt.Sprintf("%.2f", MonteCarloPi(100000000))
		t.Log(result)
		if result != "3.14" {
			t.Errorf("Wrong result! Expected:%f, returned:%s ", 3.1415, result)
		}
	})
}

func BenchmarkMonteCarloPi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonteCarloPi(100000000)
	}
}

func TestSplitInt(t *testing.T) {
	testCases := []struct {
		name           string
		x              int
		n              int
		expectedResult []int
		expectedError  bool
	}{
		{"multiple", 10, 5, []int{2, 2, 2, 2, 2}, false},
		{"n=1", 10, 1, []int{10}, false},
		{"x=10, n=3", 10, 3, []int{4, 3, 3}, false},
		{"x=10, n=7", 10, 7, []int{2, 2, 2, 1, 1, 1, 1}, false},
		{"n>x", 10, 11, nil, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := splitInt(testCase.x, testCase.n)
			for i := 0; i < len(testCase.expectedResult); i++ {
				if res[i] != testCase.expectedResult[i] {
					t.Fatalf("unexpected result at index %d: %d - expected %d", i, res[i], testCase.expectedResult[i])
				}
			}
			if testCase.expectedError {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error - %s", err)
				}
			}
		})
	}
}

func TestMonteCarloPi2(t *testing.T) {
	res, err := MonteCarloPiConcurrent(100000000)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	result := fmt.Sprintf("%.2f", res)
	t.Log(result)
	if result != "3.14" {
		t.Errorf("Wrong result! Expected:%f, returned:%s ", 3.1415, result)
	}
}

func BenchmarkMonteCarloPi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MonteCarloPiConcurrent(100000000)
		if err != nil {
			b.Fatalf("unexpected error - %s", err)
		}
	}
}
