/*
dsa.go
description: DSA encryption and decryption including key generation
details: [DSA wiki](https://en.wikipedia.org/wiki/Digital_Signature_Algorithm)
author(s): [ddaniel27](https://github.com/ddaniel27)
*/
package dsa

import (
	"crypto/rand"
	"io"
	"math/big"
)

const (
	numMRTests = 64   // Number of Miller-Rabin tests
	L          = 1024 // Number of bits in p
	N          = 160  // Number of bits in q
)

type (
	// parameters represents the DSA parameters
	parameters struct {
		P, Q, G *big.Int
	}

	// dsa represents the DSA
	dsa struct {
		parameters
		pubKey  *big.Int // public key (y)
		privKey *big.Int // private key (x)
	}
)

// New creates a new DSA instance
func New() *dsa {
	d := new(dsa)
	d.dsaParameterGeneration()
	d.keyGen()
	return d
}

// Parameter generation for DSA
// 1. FIPS 186-4 specifies that the L and N values must be (1024, 160), (2048, 224), or (3072, 256)
// 2. Choose a N-bit prime q
// 3. Choose a L-bit prime p such that p-1 is a multiple of q
// 4. Choose an integer h randomly from the range [2, p-2]
// 5. Compute g = h^((p-1)/q) mod p
// 6. Return (p, q, g)
func (dsa *dsa) dsaParameterGeneration() {
	var err error
	p, q, bigInt := new(big.Int), new(big.Int), new(big.Int)
	one, g, h := big.NewInt(1), big.NewInt(1), big.NewInt(2)
	pBytes := make([]byte, L/8)

	// GPLoop is a label for the loop
	// We use this loop to change the prime q if we don't find a prime p
GPLoop:
	for {
		// 2. Choose a N-bit prime q
		q, err = rand.Prime(rand.Reader, N)
		if err != nil {
			panic(err)
		}

		for i := 0; i < 4*L; i++ {
			// 3. Choose a L-bit prime p such that p-1 is a multiple of q
			// In this case we generate a random number of L bits
			if _, err := io.ReadFull(rand.Reader, pBytes); err != nil {
				panic(err)
			}

			// This are the minimum conditions for p being a possible prime
			pBytes[len(pBytes)-1] |= 1 // p is odd
			pBytes[0] |= 0x80          // p has the highest bit set
			p.SetBytes(pBytes)

			// Instead of using (p-1)%q == 0
			// We ensure that p-1 is a multiple of q and validates if p is prime
			bigInt.Mod(p, q)
			bigInt.Sub(bigInt, one)
			p.Sub(p, bigInt)
			if p.BitLen() < L || !p.ProbablyPrime(numMRTests) { // Check if p is prime and has L bits
				continue
			}

			dsa.P = p
			dsa.Q = q
			break GPLoop
		}
	}

	// 4. Choose an integer h randomly from the range [2, p-2]. Commonly, h = 2
	// 5. Compute g = h^((p-1)/q) mod p. In case g == 1, increment h until g != 1
	pm1 := new(big.Int).Sub(p, one)

	for g.Cmp(one) == 0 {
		g.Exp(h, new(big.Int).Div(pm1, q), p)
		h.Add(h, one)
	}

	dsa.G = g
}

// keyGen is key generation for DSA
// 1. Choose a random integer x from the range [1, q-1]
// 2. Compute y = g^x mod p
func (dsa *dsa) keyGen() {
	// 1. Choose a random integer x from the range [1, q-1]
	x, err := rand.Int(rand.Reader, new(big.Int).Sub(dsa.Q, big.NewInt(1)))
	if err != nil {
		panic(err)
	}

	dsa.privKey = x

	// 2. Compute y = g^x mod p
	dsa.pubKey = new(big.Int).Exp(dsa.G, x, dsa.P)
}

// Sign is signature generation for DSA
// 1. Choose a random integer k from the range [1, q-1]
// 2. Compute r = (g^k mod p) mod q
// 3. Compute s = (k^-1 * (H(m) + x*r)) mod q
func Sign(m []byte, p, q, g, x *big.Int) (r, s *big.Int) {
	// 1. Choose a random integer k from the range [1, q-1]
	k, err := rand.Int(rand.Reader, new(big.Int).Sub(q, big.NewInt(1)))
	if err != nil {
		panic(err)
	}

	// 2. Compute r = (g^k mod p) mod q
	r = new(big.Int).Exp(g, k, p)
	r.Mod(r, q)

	// 3. Compute s = (k^-1 * (H(m) + x*r)) mod q
	h := new(big.Int).SetBytes(m)     // This should be the hash of the message
	s = new(big.Int).ModInverse(k, q) // k^-1 mod q
	s.Mul(
		s,
		new(big.Int).Add( // (H(m) + x*r)
			h,
			new(big.Int).Mul(x, r),
		),
	)
	s.Mod(s, q) // mod q

	return r, s
}

// Verify is signature verification for DSA
// 1. Compute w = s^-1 mod q
// 2. Compute u1 = (H(m) * w) mod q
// 3. Compute u2 = (r * w) mod q
// 4. Compute v = ((g^u1 * y^u2) mod p) mod q
// 5. If v == r, the signature is valid
func Verify(m []byte, r, s, p, q, g, y *big.Int) bool {
	// 1. Compute w = s^-1 mod q
	w := new(big.Int).ModInverse(s, q)

	// 2. Compute u1 = (H(m) * w) mod q
	h := new(big.Int).SetBytes(m) // This should be the hash of the message
	u1 := new(big.Int).Mul(h, w)
	u1.Mod(u1, q)

	// 3. Compute u2 = (r * w) mod q
	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, q)

	// 4. Compute v = ((g^u1 * y^u2) mod p) mod q
	v := new(big.Int).Exp(g, u1, p)
	v.Mul(
		v,
		new(big.Int).Exp(y, u2, p),
	)
	v.Mod(v, p)
	v.Mod(v, q)

	// 5. If v == r, the signature is valid
	return v.Cmp(r) == 0
}

// GetPublicKey returns the public key (y)
func (dsa *dsa) GetPublicKey() *big.Int {
	return dsa.pubKey
}

// GetParameters returns the DSA parameters (p, q, g)
func (dsa *dsa) GetParameters() parameters {
	return dsa.parameters
}

// GetPrivateKey returns the private Key (x)
func (dsa *dsa) GetPrivateKey() *big.Int {
	return dsa.privKey
}
