package main

import "fmt"
       

func dencrypt(str string, key int)string{
	var char string
	for _, r := range str {
		newByte := int(r)
    	  if 'A' <= newByte && newByte <= 'Z'{
    	  	char += string(int('A') + int(newByte - 'A' + key) % 26)
    	  }else if 'a' <= newByte && newByte <= 'z'{
    	  	char += string(int('a') + int(newByte - 'a' + key) % 26)
    	  }else{
    	  	char += string(newByte)
    	  }
	}
	return char
}

func main(){
	fmt.Print("Enter your Message:")
	var str string
	fmt.Scan(&str)
	str_d := dencrypt(str,13)
	fmt.Println("Decryption using key => 13 is :",str_d)
	str_e := dencrypt(str_d,13)
	fmt.Print("Encryption using key => 13 is :",str_e)
}
