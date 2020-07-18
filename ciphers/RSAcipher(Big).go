package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

/*
Care has been taken to uses cryptographic secure functions
The primes numbers are 1024 bits which is as secure as u can get really
crypto/rand library has been imported as crypto and not rand
This import style will make it easier to spot all the cryptographic secure functions
*/
func main() {
	p, _ := crypto.Prime(crypto.Reader, 1024)
	q, _ := crypto.Prime(crypto.Reader, 1024)
	if !(primeCheck(p) || primeCheck(q)) {
		//they are always prime, no worries
		fmt.Println("These numbers ain't prime")
	}
	n := new(big.Int).Mul(p, q)

	one := big.NewInt(1)

	delta := lcmBig(p.Sub(p, one), q.Sub(q, one))

	e, _ := crypto.Prime(crypto.Reader, delta.BitLen())
	d := big.NewInt(0)
	d.ModInverse(e, delta)

	cleartext := "Black Lives Matter, all lives can't matter until Black lives matter"
	runes := []rune(cleartext)
	ASCIIs := toASCII(runes)
	stringEncoded := stringEncode(ASCIIs)
	bigNum, _ := new(big.Int).SetString(stringEncoded, 0)
	/*
		TODO: check that bigNum is not larger than N if larger break
		into two or more strings and encrypt separately
	*/
	fmt.Printf("Message to be encrypted: %v \n", cleartext)
	fmt.Printf("ASCII encoded: %v\n", bigNum)
	encrypted := encryptBig(bigNum, e, n)
	fmt.Printf("ciphertext: %v \n", encrypted)
	decrypted := decryptBig(encrypted, d, n)
	fmt.Printf("Decrypted but still ASCII encoded: %v \n", decrypted)
	decryptASCIIs := stringDecode(decrypted)
	fmt.Printf("Plaintext (original message) :%v", toRune(decryptASCIIs))
}

func encryptBig(num *big.Int, privateExponent *big.Int, modulus *big.Int) *big.Int {
	//encrypts by modular exponentiation
	encrypted := new(big.Int).Exp(num, privateExponent, modulus)
	return encrypted
}

func decryptBig(num *big.Int, publicExponent *big.Int, modulus *big.Int) *big.Int {
	//decrypts by modular exponentiation
	decrypted := new(big.Int).Exp(num, publicExponent, modulus)
	return decrypted
}

func lcmBig(x *big.Int, y *big.Int) *big.Int {
	//an lcm implementation for big.Int numbers
	gcd := new(big.Int).GCD(nil, nil, x, y)
	temp := new(big.Int).Mul(x, y)
	lcm := new(big.Int).Div(temp, gcd)
	return lcm
}

func primeCheck(prime *big.Int) bool {
	//primality test
	return prime.ProbablyPrime(256)
}

func toASCII(slice []rune) []int {
	//runs in O(n) where n = len(slice)
	var converted []int
	for _, v := range slice {
		converted = append(converted, int(v))
	}
	return converted
}

func toRune(slice []int) string {
	//runs in O(n) where n = len(slice)
	var str string
	for _, v := range slice {
		str += string(v)
	}
	return str
}

func stringEncode(slice []int) string {
	//encodes the ASCII to a string
	var out []string
	for _, v := range slice {
		if v < 100 {
			out = append(out, "0"+strconv.Itoa(v))
			continue
		}
		out = append(out, strconv.Itoa(v))
	}
	var str string
	for _, v := range out {
		str += v
	}
	//strips leading 0 if present to avoid conversion errors
	if str[0] == '0' {
		str = str[1:]
	}
	return str
}

func stringDecode(decryptedBig *big.Int) []int {
	//decodes the number to string then ASCII values
	str := decryptedBig.String()
	if len(str)%3 != 0 {
		str = "0" + str
	}
	var ASCII []int
	for i := 0; i < len(str); i += 3 {
		temp, _ := strconv.Atoi(str[i : i+3])
		ASCII = append(ASCII, temp)
	}
	return ASCII
}
