package main

import (
	//"math/big"
	"math/rand"
	"math"
	"time"
	"fmt"
)
func generatePrimes(limit int)int{
	/*
	generate primes by factoring
	relies on the 30k+i, though better formulae exist
	where k >=0 and i = (1,7,11,13,17,13,19,23,29)
	*/
	primes:= prime(limit)
	var choice []int
	choice = append(choice, 1,7,11,13,17,19,23,29)
	for{	
		k:=rand.Intn(int(limit/30))
		i:=choice[rand.Intn(len(choice))]
		c:=30*k+i
		found := true
		for _,v:= range primes{
			if c%v==0{
				found = false
				break
			}
		}
		if found{
			return c
		}
	}
}
func prime(limit int)(primes []int){
	sqrtLimit:=int(math.Ceil(math.Sqrt(float64(limit))))
	exit:= false
	primes = append(primes,2,3,5)
	lastIndex :=2
	for ;primes[lastIndex]<sqrtLimit;{
		if exit == true{
			break
		}
		for i:=primes[lastIndex]+2;i<primes[lastIndex]*primes[lastIndex];i+=2{
			found:= true
			for _,v:= range primes {
				if i%v==0{
					found= false
					break
				}
			}
			if found{
				primes = append(primes,i)
				lastIndex++
				if i >=sqrtLimit{
					exit = true
					break
				}
			}

		}
	}
	return
}
func lcm (a int, b int)int{
	//complexity depends on gcd
	return int((a*b)/gcd(a,b))

}
func gcd (a int, b int) int{
	//complexity not clear
	for b != 0{
		t:=b
		b = a % b
		a = t
	}
	return a
}
func modularMultiplicativeInverse(e int, delta int)int{
	//runs in O(n) where n = delta
	e= e % delta
	for i:=1;i<delta;i++{
		if (i*e)%delta==1{
			return i
		}
	}
	return 0
}

func modularExponentiation(b int, e int, mod int)int{
	//runs in O(log(n)) where n = e
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

func encryptRSA(message []int,e int,n int)[]int{
	//runs in O(k*log(n)) where k = len(message) and n = e
	var ciphertext []int
	for _,v := range message{
		ciphertext = append(ciphertext, modularExponentiation(v,e,n))
	}
	return ciphertext
}
func decryptRSA(ciphertext []int, d int, n int )[]int{
	//runs in O(k*log(n)) where k = len(ciphertext) and n = d
	var message []int
	for _,v := range ciphertext {
		message = append(message, modularExponentiation(v,d,n))
	}
	return message
}
func toASCII(slice []rune)[]int{
	//runs in O(n) where n = len(slice)
	var converted []int
	for _,v:= range slice{
		converted = append(converted, int(v))
	}
	return converted
}

func toRune(slice []int)string{
	//runs in O(n) where n = len(slice)
	var str string
	for _,v:= range slice{
		str += string(v)
	}
	return str
}
	

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	bits:=15
	
	p:= generatePrimes(1<<bits)
	q:= generatePrimes(1<<bits)
	for p==q{
		q = generatePrimes(1<<bits)
	} 
	
	n:= p*q
	
	delta:=lcm(p-1,q-1)
	
	e:=generatePrimes(delta)
	d:=modularMultiplicativeInverse(e,delta)

	fmt.Printf("%v \n%v \n%v \n%v\n",p,q,e,d)
	

	str:="I think RSA is really great"
	message := []rune(str)
	asciiSlice :=toASCII(message)

	fmt.Printf("asciiSlice :%v \n",asciiSlice)
	encrypted := encryptRSA(asciiSlice,e,n)
	fmt.Printf("encrypted :%v \n",encrypted)
	decrypted := decryptRSA(encrypted,d,n)
	fmt.Printf("decrypted : %v \n",decrypted)
	fmt.Printf("cleartext : %v \n",toRune(decrypted))
}
