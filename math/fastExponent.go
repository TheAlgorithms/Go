package main

//based on square and multiply algorithm

func fastExponentiation(b int, e int)float64{
	//runs in O(log(n)) where n = e
	if e == 0{
		return 1.0
	}
	if e == 1{
		return float64(b)
	}
	if e<0{
		e*=-1
		println("executed")
		return 1/fastExponentiation(b,e)
	}
	r:=1
	for ;e>0;{
		if e%2==1{
			r=r*b
		}
		e =e>>1
		b = b*b
	}
	return float64(r)
}

func main(){
	print(fastExponentiation(2,0))
}