package main

import "testing"

func arrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestPrimeFinder tests prime finding
func TestPrimeFind(t *testing.T) {

	if !arrayEquals(PrimesWithMax(6), []int{2, 3, 5}) {
		t.Error("[Error] Not equal PrimesWithMax(9)")
	}
}
