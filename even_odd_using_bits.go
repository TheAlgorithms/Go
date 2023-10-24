package main

import "fmt"

func isEven(n int) bool {
    // Check if the least significant bit is 0
    return n&1 == 0
}

func main() {
    number := 7

    if isEven(number) {
        fmt.Printf("%d is even.\n", number)
    } else {
        fmt.Printf("%d is odd.\n", number)
    }
}
