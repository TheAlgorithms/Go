// subset_test.go
// description: Test for generating all subsets of a set
// author(s) [sixwaaaay](https://github.com//sixwaaaay)
// see subset.go

package backtrack

import (
	"reflect"
	"testing"
)

// test Subset function
func TestSubset(t *testing.T) {
	t.Run("generate subset for given array", func(t *testing.T) {
		excepted := [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}}
		array := []int{1, 2, 3}
		if reflect.DeepEqual(excepted, Subset(array)) == false {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", excepted, Subset(array))
		}
	})
}
