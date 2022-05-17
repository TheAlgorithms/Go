package dynamic

import (
	"fmt"
	"testing"
)

func TestAbbreviation(t *testing.T) {
	aStrings := []string{
		"uOHlGMdUBc",
		"kotgDIUagj",
		"WPTffVkSNl",
		"CoJsPURrVX",
		"xasreDHndqvCnFfX",
		"XFEaWCxpeepGjOnCCsFh",
		"",
		"a",
		"a",
		"a",
		"A",
	}
	bStrings := []string{
		"uOalGMdUBCasdcavsdf",
		"DIU",
		"WPTVSN",
		"CPUVX",
		"DHndqvCnFX",
		"XFEAWCPEPGOCCSF",
		"",
		"",
		"b",
		"a",
		"A",
	}
	expected := []bool{false, true, true, false, false, true, true, true, false, false, true}
	count := len(aStrings)
	for i := 0; i < count; i++ {
		name := fmt.Sprintf(
			"Test case #%d: string a = \"%s\", string b = \"%s\"",
			i+1,
			aStrings[i],
			bStrings[i],
		)
		t.Run(name, func(t *testing.T) {
			result := Abbreviation(aStrings[i], bStrings[i])
			if result != expected[i] {
				t.Errorf("Expected the %t, got %t", expected[i], result)
			}
		})
	}
}
