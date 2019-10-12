package levenshtein

import "testing"

func TestMinPositive(t *testing.T) {
	expected := 1
	given := min(1, 2, 3)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestMinNegative(t *testing.T) {
	expected := -3
	given := min(-1, -2, -3)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestMinZero(t *testing.T) {
	expected := -1
	given := min(-1, 0, 1)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestMinAllEquals(t *testing.T) {
	expected := 42
	given := min(42, 42, 42)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestDist(t *testing.T) {
	a := "kitten"
	b := "sitting"

	expected := 3
	given := Distance(a, b, false)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestDistEquals(t *testing.T) {
	a := "kitten"
	b := "kitten"

	expected := 0
	given := Distance(a, b, false)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestDistEmptyA(t *testing.T) {
	a := ""
	b := "100-days-of-algorithms"

	expected := len(b)
	given := Distance(a, b, false)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestDistEmptyB(t *testing.T) {
	a := "Levenshtein distance"
	b := ""

	expected := len(a)
	given := Distance(a, b, false)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestDistCaseInsensitive(t *testing.T) {
	a := "Answer to the Ultimate Question of Life, the Universe, and Everything"
	b := "Answer to the ultimate question of life, the universe, and everything"

	expected := 0
	given := Distance(a, b, true)

	if given != expected {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}
