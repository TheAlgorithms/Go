package moserdebruijnsequence

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	inputNum int
	expected []int
}{
	{"first 15 terms", 15, []int{0, 1, 4, 5, 16, 17, 20, 21, 64, 65, 68, 69, 80, 81, 84}},
}

func TestMoserDeBruijnSequence(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if output := MoserDeBruijnSequence(test.inputNum); !reflect.DeepEqual(output, test.expected) {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.inputNum, test.expected, output)
			}
		})
	}
}
