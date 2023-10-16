package search

import (
	"errors"
	"testing"
)

func TestBinary(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := Binary(test.data, test.key, 0, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestBinaryIterative(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := BinaryIterative(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestLowerBound(t *testing.T) {
	for _, test := range lowerBoundTests {
		actualValue, actualError := LowerBound(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestUpperBound(t *testing.T) {
	for _, test := range upperBoundTests {
		actualValue, actualError := UpperBound(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func BenchmarkBinary(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = Binary(testCase, i, 0, len(testCase)-1)
	}
}

func BenchmarkBinaryIterative(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = BinaryIterative(testCase, i)
	}
}

func BenchmarkLowerBound(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = LowerBound(testCase, i)
	}
}

func BenchmarkUpperBound(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_, _ = UpperBound(testCase, i)
	}
}
