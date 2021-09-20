package search

import "testing"

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, _ := Binary(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, _ := BinaryIterative(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
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
		_, _ = Binary(testCase, 10, 0, len(testCase)-1)
	}
}

func BenchmarkIterBinarySearch(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = BinaryIterative(testCase, 10, 0, len(testCase)-1)
	}
}
