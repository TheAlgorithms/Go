// Package diffiehellman implements Diffie-Hellman Key Exchange Algorithm
// description: Diffie-Hellman key exchange
// details : Diffie-Hellman key exchange is a method of securely exchanging cryptographic keys over a public channel by combining private keys of two parties to generate a shared secret key.
// time complexity: O(log(n))
// space complexity: O(1)
// for more information watch : https://www.youtube.com/watch?v=NmM9HA2MQGI
package diffiehellman

const (
	generator         = 3
	primeNumber int64 = 6700417 // prime number discovered by Leonhard Euler
)

// GenerateShareKey : generates a key using client private key , generator and primeNumber
// this key can be made public
// shareKey = (g^key)%primeNumber
func GenerateShareKey(prvKey int64) int64 {
	return modularExponentiation(generator, prvKey, primeNumber)
}

// GenerateMutualKey : generates a mutual key that can be used by only alice and bob
// mutualKey = (shareKey^prvKey)%primeNumber
func GenerateMutualKey(prvKey, shareKey int64) int64 {
	return modularExponentiation(shareKey, prvKey, primeNumber)
}

// r = (b^e)%mod
func modularExponentiation(b, e, mod int64) int64 {

	//runs in O(log(n)) where n = e
	//uses exponentiation by squaring to speed up the process
	if mod == 1 {
		return 0
	}
	var r int64 = 1
	b = b % mod
	for e > 0 {
		if e&1 == 1 {
			r = (r * b) % mod
		}
		e = e >> 1
		b = (b * b) % mod
	}
	return r
}
