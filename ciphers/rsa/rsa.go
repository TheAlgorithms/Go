package rsa

import (
	"github.com/TheAlgorithms/Go/math/modulararithmetic"
	"github.com/TheAlgorithms/Go/math/sieve"
)

// Encrypt encrypts the numerical representation of the data
func Encrypt(num, publicExponent, modulus int64) (int64, error) {
	return modulararithmetic.ModularExponentiation(num, publicExponent, modulus)
}

// Decrypt decrypts the numerical representation of the encrypted data
func Decrypt(num, privateExponent, modulus int64) (int64, error) {
	return modulararithmetic.ModularExponentiation(num, privateExponent, modulus)
}

// ConvertStringToByteSlice Converts string to byte slice
func ConvertStringToByteSlice(str string) []byte {
	return []byte(str)
}

// ConvertByteSliceToIntSlice Converts byte slice to int64 slice
func ConvertByteSliceToIntSlice(slice []byte) []int64 {
	var result []int64
	for _, v := range slice {
		result = append(result, int64(v))
	}
	return result
}

// ConvertIntSliceToByteSlice Converts int64 slice to byte slice
func ConvertIntSliceToByteSlice(slice []int64) []byte {
	var result []byte
	for _, v := range slice {
		result = append(result, byte(v))
	}
	return result
}

// ConvertByteSliceToString Converts byte slice to string
func ConvertByteSliceToString(slice []byte) string {
	return string(slice)
}

// GeneratePrime Generates limit number of primes, where limit is passed as parameter
func GeneratePrime(limit int) []int {
	var primes []int

	ch := make(chan int)
	go sieve.Generate(ch)

	for i := 0; i < limit; i++ {
		primes = append(primes, <-ch)
		ch1 := make(chan int)
		go sieve.Sieve(ch, ch1, primes[i])
		ch = ch1
	}

	return primes
}

// Working implementation
// func main() {
// 	p := 61
// 	q := 53
// 	n := p * q

// 	totient := (p-1) * (q - 1)

// 	e := 17

// 	if Iterative(int64(17), int64(totient)) != 1 {
// 		panic("not coprime")
// 	}

// 	d := 2753

// 	var str = "Hello World"
// 	var b = []byte(str)

// 	fmt.Println(b)

// 	var result []int64

// 	for _, v := range b {
// 		c, err := Encrypt(int64(v), int64(e), int64(n))
// 		if err != nil {
// 			panic("something wrong")
// 		}
// 		result = append(result, c)
// 	}

// 	fmt.Println(result)

// 	var decrypted []byte
// 	for _, v := range result {
// 		c, err := Decrypt(int64(v), int64(d), int64(n))
// 		if err != nil {
// 			panic("something wrong")
// 		}
// 		decrypted = append(decrypted, byte(c))
// 	}

// 	s := string(decrypted)
// 	fmt.Println(s)
// }
