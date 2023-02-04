package math

import (
	"fmt"
	"testing"
)

func retCases() []struct {
	input        int64
	output       bool
	outputString string
} {
	return []struct {
		input        int64
		output       bool
		outputString string
	}{
		{-3112312321, false, "is not"},
		{0, false, "is not"},
		{1, true, "is"},
		{2, true, "is"},
		{109, false, "is not"},
		{145, true, "is"},
		{943, false, "is not"},
		{6327, false, "is not"},
		{40585, true, "is"},
		{9743821, false, "is not"},
		{3421488712, false, "is not"},
	}
}

func TestIsKrishnamurthyNumber(t *testing.T) {
	for _, test := range retCases() {
		t.Run(fmt.Sprintf("%d ", test.input)+test.outputString+" a Krishnamurthy Number", func(t *testing.T) {
			res := IsKrishnamurthyNumber(test.input)
			if res != test.output {
				t.Errorf("For input " + fmt.Sprintf("%d ", test.input) + "Expected: " + fmt.Sprintf("%t, ", test.output) + "Found: " + fmt.Sprintf("%t, ", res))
			}
		})
	}
}

func BenchmarkIsKrishnamurthyNumber(b *testing.B) {
	for _, test := range retCases() {
		b.Run(fmt.Sprintf("%d ", test.input)+test.outputString+" a Krishnamurthy Number", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsKrishnamurthyNumber(test.input)
			}
		})
	}
}
