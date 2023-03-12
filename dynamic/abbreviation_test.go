package dynamic

import (
	"fmt"
	"testing"
)

func TestAbbreviation(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"uOHlGMdUBc", "uOalGMdUBCasdcavsdf", false},
		{"kotgDIUagj", "DIU", true},
		{"WPTffVkSNl", "WPTVSN", true},
		{"CoJsPURrVX", "CPUVX", false},
		{"xasreDHndqvCnFfX", "DHndqvCnFX", false},
		{"XFEaWCxpeepGjOnCCsFh", "XFEAWCPEPGOCCSF", true},
		{"", "", true},
		{"a", "", true},
		{"a", "b", false},
		{"a", "a", false},
		{"A", "A", true},
	}
	count := len(tests)
	for i := 0; i < count; i++ {
		name := fmt.Sprintf(
			"Test case #%d: string a = \"%s\", string b = \"%s\"",
			i+1,
			tests[i].a,
			tests[i].b,
		)
		t.Run(name, func(t *testing.T) {
			result := Abbreviation(tests[i].a, tests[i].b)
			if result != tests[i].expected {
				t.Errorf("Expected the %t, got %t", tests[i].expected, result)
			}
		})
	}
}
