// permutation_test.go
// description: test generating permutation of a given array
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see permutation.go

package backtrack

import (
	"reflect"
	"testing"
)

// test permutation function
func TestPermutation(t *testing.T) {
	t.Run("generate permutations for given array", func(t *testing.T) {
		excepted := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}
		array := []int{1, 2, 3}
		if reflect.DeepEqual(excepted, Permutation(array)) == false {
			t.Errorf("excepted: %v, but got: %v", excepted, Permutation(array))
		}
	})
}
