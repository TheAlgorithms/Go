// factorial_test.go
// description: Test for calculating factorial
// see factorial.go

package factorial

import "testing"
import "fmt"

type factorialFun func(int) (int, error)

var implementations = map[string]factorialFun{
	"Iterative": Iterative,
	"Recursive": Recursive,
	"UsingTree": UsingTree,
}

var testCases = []struct {
	n        int
	expected int
}{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
	{6, 720},
	{7, 5040},
	{8, 40320},
	{9, 362880},
	{10, 3628800},
	{11, 39916800},
	{12, 479001600},
}

func TestFactorial(t *testing.T) {
	for implName, implFunction := range implementations {
		t.Run(implName+" errors for negative input", func(t *testing.T) {
			_, error := implFunction(-1)
			if error != ErrNegativeArgument {
				t.Errorf("No error captured for negative input")
			}
		})
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s with input %d", implName, tc.n), func(t *testing.T) {
				actual, err := implFunction(tc.n)
				if err != nil {
					t.Errorf("unexpected error captured")
				}
				if actual != tc.expected {
					t.Errorf("Expected: %d, got: %d", tc.expected, actual)
				}
			})
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for _, input := range []int{5, 10, 15} {
		for implName, implFunction := range implementations {
			b.Run(fmt.Sprintf("%s_%d", implName, input), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = implFunction(input)
				}
			})
		}
	}
}
