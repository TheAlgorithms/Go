// This program generates a password from a list of possible chars
// You must provide a minimum length and a maximum length
// This length is not fixed if you generate multiple passwords for the same range

package passwordgenerator

import (
	crand "crypto/rand"
	"io"
	"math/rand"
)

func generatePassword(minLength int, maxLength int) string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	var length = rand.Intn(maxLength-minLength) + minLength

	newPassword := make([]byte, length)
	randomData := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(crand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPassword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPassword)
			}
		}
	}
}

// func main() {
// 	rand.Seed(time.Now().Unix())

// 	fmt.Print("Please specify a minimum length: ")
// 	var minLength int
// 	fmt.Scanf("%d", &minLength)

// 	fmt.Print("Please specify a maximum length: ")
// 	var maxLength int
// 	fmt.Scanf("%d", &maxLength)

// 	fmt.Printf("Your generated password is %v\n", generatePassword(minLength, maxLength))
// }
