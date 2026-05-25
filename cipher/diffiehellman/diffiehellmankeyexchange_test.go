package diffiehellman

import (
	"crypto/rand"
	"math/big"
	"testing"
)

// randomPrivateKey returns a uniformly random private key in the range
// [1, primeNumber-1], which is the valid range for a Diffie-Hellman secret.
func randomPrivateKey(t *testing.T) int64 {
	t.Helper()
	// rand.Int returns a value in [0, primeNumber-1); shift by 1 to get [1, primeNumber-1].
	n, err := rand.Int(rand.Reader, big.NewInt(primeNumber-1))
	if err != nil {
		t.Fatalf("failed to generate random private key: %v", err)
	}
	return n.Int64() + 1
}

func TestDiffieHellmanKeyExchange(t *testing.T) {
	t.Run("Test 1: modularExponentiation", func(t *testing.T) {
		var want int64 = 9 // (3^5)mod13 = 243mod13 = 9
		var prvKey int64 = 5
		var generator int64 = 3
		var primeNumber int64 = 13
		got := modularExponentiation(generator, prvKey, primeNumber)
		if got != want {
			t.Errorf(`with privateKey=%d, generator=%d and primeNumber=%d  
			modularExponentiation should result=%d, but resulted=%d`, prvKey, generator, primeNumber, want, got)
		}
	})

	t.Run("Test 2: Key Exchange", func(t *testing.T) {
		// alice and bob each pick a private key
		alicePrvKey := randomPrivateKey(t)
		bobPrvKey := randomPrivateKey(t)

		// alice and bob generates their respective share key with their privateKey
		shareKeyByAlice := GenerateShareKey(alicePrvKey)
		shareKeyByBob := GenerateShareKey(bobPrvKey)

		// generated share key now can be exchanged even via insecure channel

		// mutualKey can be computed using shared key
		mutualKeyComputedByAlice := GenerateMutualKey(alicePrvKey, shareKeyByBob)
		mutualKeyComputedByBob := GenerateMutualKey(bobPrvKey, shareKeyByAlice)

		if mutualKeyComputedByAlice != mutualKeyComputedByBob {
			t.Errorf("mutual key computed by alice and bob should be same, but got un-equal")
		}
	})
}
