// Title: FizzBuzz
// Author: Carm1ne
// Website: thisisthecarm1ne.github.io

// 1. This program loops through numbers 1 to 100
// 2. If the number is divisible by 3 - prints "Fizz"
// 3. If the number is divisible by 5 - prints "Buzz"
// 4. If the number is divisible by both 3 and 5 - prints "FizzBuzz"
// 5. Else - prints number

package dynamic

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
