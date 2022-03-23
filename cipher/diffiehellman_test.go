package cipher_test

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/TheAlgorithms/Go/cipher"
	"testing"
)

func TestDiffieHellmanKeyExchange(t *testing.T) {
	t.Run("Key Exchange", func(t *testing.T) {
		// generating a small sized rsa_cipher key for testing
		alicePrvKey, _ := rsa.GenerateKey(rand.Reader, 31)
		bobPrvKey, _ := rsa.GenerateKey(rand.Reader, 31)

		// alice and bob generates their respective share key with their privateKey
		shareKeyByAlice, err := cipher.GenerateShareKey(alicePrvKey.D.Int64())
		if err != nil {
			t.Errorf("GenrateShareKey failed: %s", err)
		}
		shareKeyByBob, err := cipher.GenerateShareKey(bobPrvKey.D.Int64())
		if err != nil {
			t.Errorf("GenerateShareKey failed: %s", err)
		}

		// generated share key now can be exchanged even via insecure channel

		// mutualKey can be computed using shared key
		mutualKeyComputedByAlice, err := cipher.GenerateMutualKey(alicePrvKey.D.Int64(), shareKeyByBob)
		if err != nil {
			t.Errorf("GenerateMutualKey failed: %s", err)
		}
		mutualKeyComputedByBob, err := cipher.GenerateMutualKey(bobPrvKey.D.Int64(), shareKeyByAlice)
		if err != nil {
			t.Errorf("GenerateMutualKey failed: %s", err)
		}

		if mutualKeyComputedByAlice != mutualKeyComputedByBob {
			t.Errorf("mutual key computed by alice and bob should be same, but got un-equal")
		}
	})
}
