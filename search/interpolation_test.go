package search

import "testing"

func TestInterpolation(t *testing.T) {
	for _, test := range searchTests {
		actual := Interpolation(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test %s failed", test.name)
		}
	}
}
