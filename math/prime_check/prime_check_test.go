package main

import (
	"testing"
)

func TestTableNaiveApproach(t *testing.T) {
	var tests = []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{1, false},
		{10, false},
		{23, true},
	}

	for _, test := range tests {
		if output := NaiveApproach(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}
func TestTablePairApproach(t *testing.T) {
	var tests = []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{1, false},
		{10, false},
		{23, true},
	}

	for _, test := range tests {
		if output := PairApproach(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}
