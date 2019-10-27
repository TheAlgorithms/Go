

package main

import "fmt"

func fib() func() int {
   f1, f2 := 0, 1
   return func() int {
	f1, f2 = f2, f1+f2
	   return f1
   }
}

func main() {
   f := fib()
   for i := 0; i < 10; i++ {
	fmt.Println(f())
   }
}
