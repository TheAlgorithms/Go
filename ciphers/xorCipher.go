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

func main(){
	//string = hello world
	key :=97
	message:=[]int{72,101,108,108,111,32,119,111,114,108,100}
	encrypted:=encrypt(key,message)
	decrypted:=decrypt(key,encrypted)
	plaintext:=decodeToString(decrypted)
	fmt.Println(plaintext)
}

