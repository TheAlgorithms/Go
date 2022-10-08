// twin_test.go
// description: Returns Twin Prime of n
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see twin.go

package prime

import (
	"testing"
)

func TestTwin(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue int
	}{
		{"n = 3, should return 5", 3, 5},
		{"n = 4, should return -1", 4, -1},
		{"n = 5, should return 7", 5, 7},
		{"n = 17, should return 19", 17, 19},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Twin(test.n)
			if result != test.expectedValue {
				t.Errorf("expected value: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}
func BenchmarkTwin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Twin(65536)
	}
}
