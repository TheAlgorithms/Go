package sort

import (
	"reflect"
	"testing"
)

func TestTournament(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"unsorted", []int{5, 3, 1, 4, 2}, []int{1, 2, 3, 4, 5}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"single element", []int{42}, []int{42}},
		{"empty", []int{}, []int{}},
		{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"negative numbers", []int{-3, 0, -1, 5, 2}, []int{-3, -1, 0, 2, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tournament(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Tournament(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
