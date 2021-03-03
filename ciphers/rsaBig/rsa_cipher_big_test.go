package bigrsacipher

import (
	crypto "crypto/rand"
	"math/big"
	"testing"
)

var rsaTestData = []struct {
	description string
	input       string
}{
	{
		"Encrypt letter 'a'",
		"a",
	},
	{
		"Encrypt 'hello world'",
		"hello world",
	},
	{
		"Encrypt full sentence",
		"the quick brown fox jumps over the lazy dog.",
	},
	{
		"Encrypt full sentence from RSAcipher(Big).go main function",
		"Black Lives Matter, all lives can't matter until Black lives matter",
	},
}

func TestRSABigCipher(t *testing.T) {
	p, _ := crypto.Prime(crypto.Reader, 1024)
	q, _ := crypto.Prime(crypto.Reader, 1024)
	n := new(big.Int).Mul(p, q)

	one := big.NewInt(1)

	delta := lcmBig(p.Sub(p, one), q.Sub(q, one))

	e, _ := crypto.Prime(crypto.Reader, delta.BitLen())
	d := big.NewInt(0)
	d.ModInverse(e, delta)

	for _, test := range rsaTestData {
		t.Run(test.description, func(t *testing.T) {

			runes := []rune(test.input)
			ASCIIs := toASCII(runes)
			stringEncoded := stringEncode(ASCIIs)
			bigNum, _ := new(big.Int).SetString(stringEncoded, 0)

			encrypted := EncryptBig(bigNum, e, n)
			decrypted := DecryptBig(encrypted, d, n)
			decryptASCIIs := stringDecode(decrypted)

			if actual := ToRune(decryptASCIIs); actual != test.input {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.input, actual)
			}
		})
	}
}
