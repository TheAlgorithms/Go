package searches

import "testing"

type searchTest struct {
	data     []int
	key      int
	expected int
	name     string
}

var searchTests = []searchTest{
	//Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, "Empty"},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := BinarySearch(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actual)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := IterBinarySearch(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actual)
		}
	}
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := LinearSearch(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actual)
		}
	}
}
