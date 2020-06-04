package main

import "fmt"
func encrypt(key int, plaintext []int)(ciphertext []int){
	for _,char:= range plaintext {
		ciphertext = append(ciphertext,xor(char,key))
	}
	return
}
func xor(char int,key int)int{
	return (char^key)
}

func decrypt(key int, ciphertext []int)(plaintext []int){
	for _,char:= range ciphertext {
		plaintext = append(plaintext , xor(char,key))
	}
	return
}

func decodeToString(slice []int)(str string){
	for _,v:= range slice{
		str+=string(v)
	}
	return
}
func toASCII(slice []rune)[]int{
	var converted []int
	for _,v:= range slice{
		converted = append(converted, int(v))
	}
	return converted
}
func main(){
	str := "hello world"
	key :=97
	temp:= []rune(str)
	message:=toASCII(temp)
	encrypted:=encrypt(key,message)
	decrypted:=decrypt(key,encrypted)
	plaintext:=decodeToString(decrypted)
	fmt.Println(plaintext)
}

