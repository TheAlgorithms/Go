package cipher

import (
	"github.com/TheAlgorithms/Go/math/modular"
)

// RSA encrypts based on the RSA algorithm - uses modular exponentitation in math directory
func RSA(message []rune, publicExponent, modulus int64) ([]rune, error) {
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), publicExponent, modulus)
		if err != nil {
			return nil, ErrorFailedToEncrypt
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}

	return encrypted, nil
}
