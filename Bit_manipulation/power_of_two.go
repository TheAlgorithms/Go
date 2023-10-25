//To check whether a number is power of two
package main

import "fmt"

func isPowerOfTwo(n int) bool {
    // A number is a power of two if and only if it has exactly one '1' bit in its binary representation.
    // Using n & (n-1) will turn off the rightmost '1' bit.
    return n > 0 && (n&(n-1)) == 0
}

func main() {
    num := 16

    if isPowerOfTwo(num) {
        fmt.Printf("%d is a power of two.\n", num)
    } else {
        fmt.Printf("%d is not a power of two.\n", num)
    }
}
