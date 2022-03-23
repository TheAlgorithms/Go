package cipher

import "github.com/TheAlgorithms/Go/math/modular"

const (
	generator         = 3
	primeNumber int64 = 6700417 // prime number discovered by Leonhard Euler
)

// GenerateShareKey : generates a key using client private key , generator and primeNumber
// this key can be made public
// shareKey = (g^key)%primeNumber
func GenerateShareKey(prvKey int64) (int64, error) {
	return modular.Exponentiation(generator, prvKey, primeNumber)
}

// GenerateMutualKey : generates a mutual key that can be used by only alice and bob
// mutualKey = (shareKey^prvKey)%primeNumber
func GenerateMutualKey(prvKey, shareKey int64) (int64, error) {
	return modular.Exponentiation(shareKey, prvKey, primeNumber)
}
