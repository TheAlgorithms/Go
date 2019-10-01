package main
import "fmt"
func encryptdecrypt(input, key string) (output string) {
	for i := 0;i<len(input); i++{
		output+=string(input[i] ^ key[i % len(key)])
	}
	return output
}
func main(){
	key:="KEY"
	data:="Sample Data"

	encrypted_string := encryptdecrypt(data)
	fmt.Println("Encrypted:",encrypted_string)

	decrypted_string := encryptdecrypt(data)
	fmt.Println("Decrypted:",decrypted_string)

}