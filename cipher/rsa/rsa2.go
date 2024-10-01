/*
rsa2.go
description: RSA encryption and decryption including key generation
details: [RSA wiki](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
time complexity: O(n)
space complexity: O(1)
author(s): [ddaniel27](https://github.com/ddaniel27)
*/
package rsa

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/TheAlgorithms/Go/math/gcd"
	"github.com/TheAlgorithms/Go/math/lcm"
	"github.com/TheAlgorithms/Go/math/modular"
	"github.com/TheAlgorithms/Go/math/prime"
)

// rsa struct contains the public key, private key and modulus
type rsa struct {
	publicKey  uint64
	privateKey uint64
	modulus    uint64
}

// New initializes the RSA algorithm
// returns the RSA object
func New() *rsa {
	// The following code generates keys for RSA encryption/decryption
	// 1. Choose two large prime numbers, p and q and compute n = p * q
	p, q := randomPrime() // p and q stands for prime numbers
	modulus := p * q      // n stands for common number

	// 2. Compute the totient of n, lcm(p-1, q-1)
	totient := uint64(lcm.Lcm(int64(p-1), int64(q-1)))

	// 3. Choose an integer e such that 1 < e < totient(n) and gcd(e, totient(n)) = 1
	publicKey := uint64(2) // e stands for encryption key (public key)
	for publicKey < totient {
		if gcd.Recursive(int64(publicKey), int64(totient)) == 1 {
			break
		}
		publicKey++
	}

	// 4. Compute d such that d * e ≡ 1 (mod totient(n))
	inv, _ := modular.Inverse(int64(publicKey), int64(totient))
	privateKey := uint64(inv)

	return &rsa{
		publicKey:  publicKey,
		privateKey: privateKey,
		modulus:    modulus,
	}
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

// randomPrime returns two random prime numbers
func randomPrime() (uint64, uint64) {
	sieve := prime.SieveEratosthenes(1000)
	sieve = sieve[10:] // remove first 10 prime numbers (small numbers)
	index1 := rand.Intn(len(sieve))
	index2 := rand.Intn(len(sieve))

	for index1 == index2 {
		index2 = rand.Intn(len(sieve))
	}

	return uint64(sieve[index1]), uint64(sieve[index2])
}
