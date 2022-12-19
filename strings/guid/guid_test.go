package guid

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("check for allowed characters", func(t *testing.T) {
		allowedChars := "0123456789abcdef-"
		guid, err := New()
		if err != nil {
			t.Errorf(`the test failed, an error occurred: %s`, err.Error())
		}

		for _, char := range guid {
			if !strings.Contains(allowedChars, string(char)) {
				t.Errorf(`allowed only "%s" characters, but got %v`, allowedChars, char)
			}
		}
	})

	t.Run("check string length", func(t *testing.T) {
		guid, err := New()
		if err != nil {
			t.Errorf(`the test failed, an error occurred: %s`, err.Error())
		}

		if len(guid) != len(pattern) {
			t.Errorf(`the length of the string should be "%d", but got %d`, len(pattern), len(guid))
		}
	})

	t.Run("check for version index", func(t *testing.T) {
		expected := "4"
		versionIndex := strings.Index(pattern, expected)
		guid, err := New()
		if err != nil {
			t.Errorf(`the test failed, an error occurred: %s`, err.Error())
		}

		result := string(guid[versionIndex])

		if expected != result {
			t.Errorf(`at the index %d should be %s, but got %s`, versionIndex, expected, result)
		}
	})

	t.Run("check the number of dashes", func(t *testing.T) {
		expected := strings.Count(pattern, "-")
		guid, err := New()
		if err != nil {
			t.Errorf(`the test failed, an error occurred: %s`, err.Error())
		}

		result := strings.Count(guid, "-")

		if expected != result {
			t.Errorf(`the length of the string should be "%d", but got %d`, len(pattern), len(guid))
		}
	})
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = New()
	}
}
