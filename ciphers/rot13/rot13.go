package rot13

import (
	"bytes"
	"strings"
)

func rot13(input string) string {
	var outputBuffer bytes.Buffer
	for _, r := range strings.ToLower(input) {
		newByte := int(r)

		if newByte >= 'a' && newByte <= 'z' {
			newByte += 13

			if newByte > 'z' {
				newByte -= 26
			} else if newByte < 'a' {
				newByte += 26
			}
		}

		outputBuffer.WriteString(string(newByte))
	}
	return outputBuffer.String()
}
