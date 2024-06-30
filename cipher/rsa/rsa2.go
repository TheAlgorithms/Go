package rsa

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/TheAlgorithms/Go/math/gcd"
)

type rsa struct {
	publicKey  uint64
	privateKey uint64
	modulus    uint64
}

// InitRSA initializes the RSA algorithm
// returns the RSA object
func InitRSA() *rsa {
	rsa := new(rsa)
	rsa.keyGen()
	return rsa
}

// KeyGen generates a key for encryption and decryption with RSA
// 1. Choose two large prime numbers, p and q and compute n = p * q
// 2. Compute the totient of n, lcm(p-1, q-1) = (p-1) * (q-1)
// 3. Choose an integer e such that 1 < e < lcm(p-1, q-1) and gcd(e, lcm(p-1, q-1)) = 1
// 4. Compute d such that d * e ≡ 1 (mod lcm(p-1, q-1))
func (rsa *rsa) keyGen() {
	// 1. Choose two large prime numbers, p and q and compute n = p * q
	p, q := uint64(173), uint64(479) // prime numbers
	modulus := p * q                 // n stands for common number

	// 2. Compute the totient of n, lcm(p-1, q-1) = (p-1) * (q-1)
	totient := (p - 1) * (q - 1)

	// 3. Choose an integer e such that 1 < e < totient(n) and gcd(e, totient(n)) = 1
	publicKey := uint64(2) // e stands for encryption key (public key)
	for publicKey < totient {
		if gcd.Recursive(int64(publicKey), int64(totient)) == 1 {
			break
		}
		publicKey++
	}

	// 4. Compute d such that d * e ≡ 1 (mod totient(n))
	// d*e = 1 + k * totient
	// d = (1 + k * totient) / e
	k := uint64(5)                                // constant
	privateKey := (1 + (k * totient)) / publicKey // d stands for decryption key (private key)

	rsa.publicKey = publicKey
	rsa.privateKey = privateKey
	rsa.modulus = modulus
}

// EncryptString encrypts the data using RSA algorithm
// returns the encrypted string
func (rsa *rsa) EncryptString(data string) string {
	var nums []byte

	for _, char := range data {
		slice := make([]byte, 8)
		binary.BigEndian.PutUint64( // convert uint64 to byte slice
			slice,
			encryptDecryptInt(rsa.publicKey, rsa.modulus, uint64(char)), // encrypt each character
		)
		nums = append(nums, slice...)
	}

	return string(nums)
}

// DecryptString decrypts the data using RSA algorithm
// returns the decrypted string
func (rsa *rsa) DecryptString(data string) string {
	result := ""
	middle := []byte(data)

	for i := 0; i < len(middle); i += 8 {
		if i+8 > len(middle) {
			break
		}

		slice := middle[i : i+8]
		num := binary.BigEndian.Uint64(slice) // convert byte slice to uint64
		result += fmt.Sprintf("%c", encryptDecryptInt(rsa.privateKey, rsa.modulus, num))
	}

	return result
}

// GetPublicKey returns the public key and modulus
func (rsa *rsa) GetPublicKey() (uint64, uint64) {
	return rsa.publicKey, rsa.modulus
}

// GetPrivateKey returns the private key
func (rsa *rsa) GetPrivateKey() uint64 {
	return rsa.privateKey
}

// encryptDecryptInt encrypts or decrypts the data using RSA algorithm
func encryptDecryptInt(e, n, data uint64) uint64 {
	pow := new(big.Int).Exp(big.NewInt(int64(data)), big.NewInt(int64(e)), big.NewInt(int64(n)))
	return pow.Uint64()
}