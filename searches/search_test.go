package main

import "testing"

type searchTest struct {
	data     []int
	key      int
	expected int
	name     string
}

var searchTests = []searchTest{
	//Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, "Sanity"},
	//First
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "First"},
	//Last
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "Last"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, "Empty"},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := binarySearch(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test %s failed: expected '%d', get '%d'", test.name, test.expected, actual)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := iterBinarySearch(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test %s failed: expected '%d', get '%d'", test.name, test.expected, actual)
		}
	}
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := linearSearch(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test %s failed: expected '%d', get '%d'", test.name, test.expected, actual)
		}
	}
}
