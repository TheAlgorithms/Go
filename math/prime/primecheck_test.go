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

func TestMillerRabinTest(t *testing.T) {
	var tests = []struct {
		name     string
		input    int64
		expected bool
		rounds   int64
		err      error
	}{
		{"smallest prime", 2, true, 5, nil},
		{"random prime", 3, true, 5, nil},
		{"neither prime nor composite", 1, false, 5, nil},
		{"random non-prime", 10, false, 5, nil},
		{"another random prime", 23, true, 5, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := MillerRabinTest(test.input, test.rounds)
			if err != test.err {
				t.Errorf("For input: %d, unexpected error: %v, expected error: %v", test.input, err, test.err)
			}
			if output != test.expected {
				t.Errorf("For input: %d, expected: %v", test.input, output)
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
