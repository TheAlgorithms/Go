// rsa.go
// description: Simple RSA algorithm implementation
// details:
// A simple RSA Encryption and Decryption algorithm.
// It uses prime numbers that fit in int64 datatypes and
// thus both the Encrypt and Decrypt are not a production
// ready implementation. The OpenSSL implementation of RSA
// also adds a padding which is not present in this algorithm.
// time complexity: O(n)
// space complexity: O(n)
// author(s) [Taj](https://github.com/tjgurwara99)
// see rsa_test.go

// Package rsa shows a simple implementation of RSA algorithm
package rsa

import (
	"errors"

	modular "github.com/TheAlgorithms/Go/math/modular"
)

// ErrorFailedToEncrypt Raised when Encrypt function fails to encrypt the message
var ErrorFailedToEncrypt = errors.New("failed to Encrypt")

// ErrorFailedToDecrypt Raised when Decrypt function fails to decrypt the encrypted message
var ErrorFailedToDecrypt = errors.New("failed to Decrypt")

// Encrypt encrypts based on the RSA algorithm - uses modular exponentitation in math directory
func Encrypt(message []rune, publicExponent, modulus int64) ([]rune, error) {
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

// Decrypt decrypts encrypted rune slice based on the RSA algorithm
func Decrypt(encrypted []rune, privateExponent, modulus int64) (string, error) {
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedLetter, err := modular.Exponentiation(int64(letter), privateExponent, modulus)
		if err != nil {
			return "", ErrorFailedToDecrypt
		}
		decrypted = append(decrypted, rune(decryptedLetter))
	}
	return string(decrypted), nil
}
