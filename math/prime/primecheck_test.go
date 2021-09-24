package prime

import (
	"testing"
)

func TestTableNaiveApproach(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected bool
	}{
		{"smallest prime", 2, true},
		{"random prime", 3, true},
		{"neither prime nor composite", 1, false},
		{"random non-prime", 10, false},
		{"another random prime", 23, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := NaiveApproach(test.input); output != test.expected {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}

}
func TestTablePairApproach(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected bool
	}{
		{"smallest prime", 2, true},
		{"random prime", 3, true},
		{"neither prime nor composite", 1, false},
		{"random non-prime", 10, false},
		{"another random prime", 23, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := NaiveApproach(test.input); output != test.expected {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}

}

func BenchmarkNaiveApproach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveApproach(23)
	}
}

func BenchmarkPairApproach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PairApproach(23)
	}
}
