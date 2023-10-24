package main

import "fmt"

func power(base, exponent int) int {
    result := 1

    for exponent > 0 {
        // If the least significant bit is 1, multiply the result by base
        if exponent&1 != 0 {
            result *= base
        }

        // Shift bits to the right to consider the next bit
        exponent >>= 1

        // Square the base (equivalent to raising it to the power of 2)
        base *= base
    }

    return result
}

func main() {
    baseNumber := 2
    exponentValue := 5
    result := power(baseNumber, exponentValue)
    fmt.Printf("%d^%d is: %d\n", baseNumber, exponentValue, result)
}
