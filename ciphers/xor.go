package main

import (
	"crypto/rand"
	"math/big"
)

func main() {
	message := "Hello World!"
	key := genKey(len(message))
	em := encrypt(message, key)
	dm := decrypt(em, key)

	println(dm)
}

func encrypt(message string, key string) []byte{

	// Convert inputs to byte arrays
	binMessage := []byte(message)
	binKey := []byte(key)
	result := make([]byte, len(binMessage))

	// xor the message and the key
	for i := 0; i < len(binMessage); i++ {
		result[i] = binMessage[i] ^ binKey[i]
	}

	return result
}

func decrypt(message []byte, key string)  string{
	// Convert inputs to byte arrays
	binKey := []byte(key)
	result := make([]byte, len(message))

	// xor the message and the key
	for i := 0; i < len(message); i++ {
		result[i] = message[i] ^ binKey[i]
	}

	return string(result)
}


func genKey(length int) string{
	result := ""
	for {
		if len(result) >= length {
			return result
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return ""
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 32 && n < 127 {
			result += string(n)
		}
	}
}