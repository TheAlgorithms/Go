package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := Binary(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test %s failed", test.name)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual := BinaryIterative(test.data, test.key, 0, len(test.data)-1)
		if actual != test.expected {
			t.Errorf("test %s failed", test.name)
		}
	}
}
