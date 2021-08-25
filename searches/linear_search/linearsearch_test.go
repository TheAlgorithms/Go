package linear_search

import (
	"testing"
)

type searchTest struct {
	data     []int
	key      int
	expected int
	name     string
}

var searchTests = []searchTest{
	//Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, "Empty"},
}


func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := LinearSearch(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test %s failed", test.name)
		}
	}
}
