package main

import (
	"fmt"
	"math/big"
)

func fib(n int) *big.Int {
	fn := make([]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fib(300))
}
