package genetic

import (
	"testing"
)

func TestSimple(t *testing.T) {
	target := "This is a genetic algorithm to evaluate, combine, evolve and mutate a string!"
	charmap := []rune(" ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,;!?+-*#@^'èéòà€ù=)(&%$£/\\")

	res, err := GeneticString(target, charmap, &Conf{})
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	got := res.Best

	if got.Value != float64(1.0) {
		t.Errorf("Target value not reached\nwant: %f\n, got: %f\n", float64(1.0), got.Value)
	}

	if got.Key != target {
		t.Errorf("Target string not reached\nwant: %s\n, got: %s\n", target, got.Key)
	}
}
