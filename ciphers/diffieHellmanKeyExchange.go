package main

import (
	"fmt"
	"math/rand"
)
func main(){
	bit:=30
	/*
	p and g are pre-agreed constants
	that can be communicated over an insecure channel
	p should ideally be a large prime number but any integer works
	g should be a small integer, 2,3 works fine

	PS: Note that the secret keys are never send over
	the network
	*/

	p:=2+rand.Intn(1<<bit)
	g:=2+rand.Intn(5)

	//Both parties choose a secret key

	AliceSecret := 1 + rand.Intn(1<<bit)
	BobSecret := 1 + rand.Intn(1<<bit)

	fmt.Printf("Alice's secret key is: %v\n",AliceSecret)
	fmt.Printf("Bob's secret key is: %v\n",BobSecret)

	//Both parties send ((g^secret_key)%p)
	//It's not possible to determine the secretkey from the value sent

	AliceSends :=modularExponentiation(g,AliceSecret,p)
	BobSends :=modularExponentiation(g,BobSecret,p)

	fmt.Printf("Alice sends to Bob: %v\n",AliceSends)
	fmt.Printf("Bob sends to Alice: %v\n",BobSends)

	//Both parties calculate the shared secret key from the value send
	//(value_sent^secret_key)%p
	//Both calculations end up with same value despite the different inputs
	AliceComputes :=modularExponentiation(BobSends,AliceSecret,p)
	BobComputes := modularExponentiation(AliceSends,BobSecret,p)

	fmt.Printf("Alice Computes the shared secret key as: %v\n",AliceComputes)
	fmt.Printf("Bob Computes the shared secret key as: %v\n",BobComputes)

	// simply confirms that the values are equal
	if AliceComputes == BobComputes{
		sharedKey:=AliceComputes
		fmt.Println("Voila, shared key is ",sharedKey)
	}




}
func modularExponentiation(b int, e int, mod int)int{
	//runs in O(log(n)) where n = e
	//uses exponentiation by squaring to speed up the process
	if mod == 1{
		return 0
	}
	r:=1
	b = b % mod
	for ;e>0;{
		if e%2==1{
			r=(r*b)%mod
		}
		e =e>>1
		b = (b*b)%mod
	}
	return r
}
