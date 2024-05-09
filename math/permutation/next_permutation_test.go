package permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var nextPermutationTestData = []struct {
		description string
		numbers     []int
		next        []int
	}{
		{
			description: "Basic case",
			numbers:     []int{1, 2, 3},
			next:        []int{1, 3, 2},
		},
		{
			description: "Should reverse the whole slice",
			numbers:     []int{3, 2, 1},
			next:        []int{1, 2, 3},
		},
		{
			description: "A more complex test",
			numbers:     []int{2, 4, 1, 7, 5, 0},
			next:        []int{2, 4, 5, 0, 1, 7},
		},
	}
	for _, test := range nextPermutationTestData {
		t.Run(test.description, func(t *testing.T) {
			NextPermutation(test.numbers)

			if !reflect.DeepEqual(test.numbers, test.next) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expected result:%v\nFound: %v", test.next, test.numbers)
			}
		})
	}
}
