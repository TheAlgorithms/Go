package main

import (
	"bytes"
	"fmt"
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
func main() {
	cleartext := "We'll just make him an offer he can't refuse... tell me you get the pop culture reference"
	encrypted := rot13(cleartext)
	decrypted := rot13(encrypted)
	fmt.Printf("Cleartext: %v\nencrypted: %v\ndecrypted: %v\n", cleartext, encrypted, decrypted)
}
