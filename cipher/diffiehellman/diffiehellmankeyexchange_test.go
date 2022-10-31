package diffiehellman

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

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
		// generating a small sized rsa_cipher key for testing
		alicePrvKey, _ := rsa.GenerateKey(rand.Reader, 31)
		bobPrvKey, _ := rsa.GenerateKey(rand.Reader, 31)

		// alice and bob generates their respective share key with their privateKey
		shareKeyByAlice := GenerateShareKey(alicePrvKey.D.Int64())
		shareKeyByBob := GenerateShareKey(bobPrvKey.D.Int64())

		// generated share key now can be exchanged even via insecure channel

		// mutualKey can be computed using shared key
		mutualKeyComputedByAlice := GenerateMutualKey(alicePrvKey.D.Int64(), shareKeyByBob)
		mutualKeyComputedByBob := GenerateMutualKey(bobPrvKey.D.Int64(), shareKeyByAlice)

		if mutualKeyComputedByAlice != mutualKeyComputedByBob {
			t.Errorf("mutual key computed by alice and bob should be same, but got un-equal")
		}
	})
}

func FuzzGenSharedKey(f *testing.F) {
	f.Add(int64(13))
	f.Fuzz(func(t *testing.T, prvKey int64) {
		// explicitly ignoring the return value
		// here we want to test that no strange behaviors are raised when executing GenerateShareKey
		_ = GenerateShareKey(prvKey)
	})
}

func FuzzGenMutualKey(f *testing.F) {
	f.Add(int64(5), int64(17))
	f.Fuzz(func(t *testing.T, prvKey, shrdKey int64) {
		// explicitly ignoring the return value
		// here we want to test that no strange behaviors are raised when executing GenerateMutualKey
		_ = GenerateMutualKey(prvKey, shrdKey)
	})
}
