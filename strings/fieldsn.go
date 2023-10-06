// fields.go
// description: An algorithm which splits a string into n+1 fields.
// author(s) [Veer](https://github.com/celestix)
// see fields_test.go
package strings

import "strings"

// FieldsN splits a string into n fields.
// It is similar to strings.Fields, but it splits the string into n fields.
// If the string has less than n fields, the last field will contain the rest of the string.
// If the string has more than n fields, the last field will contain the rest of the string.
// If n is -1, it will behave like strings.Fields.
func FieldsN(s string, n int) []string {
	if n < 0 {
		return strings.Fields(s)
	}
	a := make([]string, n+1)
	b := strings.Builder{}
	index := 0
	for _, c := range s {
		if index+1 > n {
			b.WriteRune(c)
			continue
		}
		switch c {
		case '\n', '\t', '\f', '\r', ' ':
			a[index] = b.String()
			index++
			b.Reset()
			continue
		}
		b.WriteRune(c)
	}
	a[index] = b.String()
	return a[:index+1]
}
