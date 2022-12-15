// guid.go
// description: Generate random globally unique identifiers (GUIDs).
// details:
// A GUID (globally unique identifier) is a 128-bit text string that
// represents an identification (ID). Organizations generate GUIDs when
// a unique reference number is needed to identify information on
// a computer or network. A GUID can be used to ID hardware, software,
// accounts, documents and other items. The term is also often used in
// software created by Microsoft.
// See more information on: https://en.wikipedia.org/wiki/Universally_unique_identifier
// author(s) [cheatsnake](https://github.com/cheatsnake)
// see guid_test.go

// Package guid provides facilities for generating random globally unique identifiers.
package guid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const pattern string = "xxxxxxxx-xxxx-4xxx-xxxx-xxxxxxxxxxxx"

// New returns a randomly generated global unique identifier.
func New() (string, error) {
	var guid strings.Builder

	for i := range pattern {
		char := string(pattern[i])

		if char == "x" {
			random, err := rand.Int(rand.Reader, big.NewInt(16))
			if err != nil {
				return "", err
			}

			char = fmt.Sprintf("%x", random.Int64())
		}

		guid.WriteString(char)
	}

	return guid.String(), nil
}
