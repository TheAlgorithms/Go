// This program generates a password from a list of possible chars
// You must provide a minimum length and a maximum length
// This length is not fixed if you generate multiple passwords for the same range

package password_generator

import (
	"crypto/rand"
	"io"
	"math/big"
)

// GeneratePassword returns a newly generated password
func GeneratePassword(minLength int, maxLength int) string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength)))
	if err != nil {
		panic(err) // handle this gracefully
	}
	length.Add(length, big.NewInt(int64(minLength)))

	intLength := int(length.Int64())

	newPassword := make([]byte, intLength)
	randomData := make([]byte, intLength+intLength/4)
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPassword[i] = chars[c%clen]
			i++
			if i == intLength {
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
