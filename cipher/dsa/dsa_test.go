package dsa_test

import (
	"math/big"
	"testing"

	"github.com/TheAlgorithms/Go/cipher/dsa"
)

func TestDSA(t *testing.T) {
	tests := []struct {
		name    string
		message string
		alter   bool
		want    bool
	}{
		{
			name:    "valid signature",
			message: "Hello, world!",
			alter:   false,
			want:    true,
		},
		{
			name:    "invalid signature",
			message: "Hello, world!",
			alter:   true,
			want:    false,
		},
	}
	// Generate a DSA key pair
	dsaInstance := dsa.New()
	pubKey := dsaInstance.GetPublicKey()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Sign the message
			r, s := dsa.Sign(
				[]byte(tt.message),
				params.P,
				params.Q,
				params.G,
				privKey,
			)

			if tt.alter {
				// Alter the signature
				r.Add(r, big.NewInt(1))
			}

			// Verify the signature
			if got := dsa.Verify(
				[]byte(tt.message),
				r,
				s,
				params.P,
				params.Q,
				params.G,
				pubKey,
			); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

/* ------------------- BENCHMARKS ------------------- */
func BenchmarkDSANew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dsa.New()
	}
}

func BenchmarkDSASign(b *testing.B) {
	dsaInstance := dsa.New()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()
	for i := 0; i < b.N; i++ {
		dsa.Sign(
			[]byte("Hello, World!"),
			params.P,
			params.Q,
			params.G,
			privKey,
		)
	}
}

func BenchmarkDSAVerify(b *testing.B) {
	dsaInstance := dsa.New()
	pubKey := dsaInstance.GetPublicKey()
	params := dsaInstance.GetParameters()
	privKey := dsaInstance.GetPrivateKey()
	r, s := dsa.Sign(
		[]byte("Hello, World!"),
		params.P,
		params.Q,
		params.G,
		privKey,
	)
	for i := 0; i < b.N; i++ {
		dsa.Verify(
			[]byte("Hello, World!"),
			r,
			s,
			params.P,
			params.Q,
			params.G,
			pubKey,
		)
	}
}
