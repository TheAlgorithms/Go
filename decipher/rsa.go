package decipher

import (
	"errors"
	"github.com/TheAlgorithms/Go/math/modular"
)

// RSA decrypts encrypted rune slice based on the RSA algorithm
func RSA(encrypted []rune, privateExponent, modulus int64) (string, error) {
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedLetter, err := modular.Exponentiation(int64(letter), privateExponent, modulus)
		if err != nil {
			return "", ErrorFailedToDecipher
		}
		decrypted = append(decrypted, rune(decryptedLetter))
	}
	return string(decrypted), nil
}

// ErrorFailedToDecipher Raised a function fails to decrypt the encrypted message
var ErrorFailedToDecipher = errors.New("failed to Decrypt")
