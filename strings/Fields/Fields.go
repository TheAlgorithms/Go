/*
This algorithm splits the string into upto the 1st N numbers of elements separated by spaces, newlines, etc.

Example:
using FieldsN("This is a string", 1) will separate the string into []string{"This", "is a string"} 
*/

package Fields

import "strings"

func FieldsN(s string, n int) []string {
	a := make([]string, 0)
	b := strings.Fields(s)
	for i := 1; i <= n; i++ {
		if i > len(b)-1 {
			continue
		}
		a = append(a, b[i-1])
		for _, x := range []string{"\t", "\n", "\f", "\r", " ", ""} {
			s = strings.Replace(s, b[i-1]+x, "", 1)
			s = strings.Replace(s, x+b[i-1], "", 1)
		}
	}
	a = append(a, s)
	return a
}
