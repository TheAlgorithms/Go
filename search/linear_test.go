package search

import (
	"testing"
)

func TestLinear(t *testing.T) {
	for _, test := range searchTests {
		actual := Linear(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test %s failed", test.name)
		}
	}
}
