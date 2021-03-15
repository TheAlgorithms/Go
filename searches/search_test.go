package searches

import (
	"testing"
)

type searchTest struct {
	data          []int
	key           int
	expectedValue int
	expectedError error
	name          string
}

var searchTests = []searchTest{
	//Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, errNotFound, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, errNotFound, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, errNotFound, "Empty"},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := BinarySearch(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expectedValue {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expectedValue, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := IterBinarySearch(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expectedValue {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expectedValue, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := LinearSearch(test.data, test.key)
		if actual != test.expectedValue {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expectedValue, actual)
		}
	}
}

func generateBenchmarkTestCase() []int {
	var testCase []int
	for i := 0; i < 100; i++ {
		testCase = append(testCase, i)
	}
	return testCase
}

func BenchmarkBinarySearch(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearch(testCase, 10, 0, len(testCase)-1)
	}
}

func BenchmarkIterBinarySearch(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = IterBinarySearch(testCase, 10, 0, len(testCase)-1)
	}
}
