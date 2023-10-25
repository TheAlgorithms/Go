//Check for Duplicates in a List i.e,[2, 3, 1, 4, 2, 3, 1]=>4
package main

import "fmt"

func findSingle(nums []int) int {
    result := 0

    for _, num := range nums {
        result ^= num
    }

    return result
}

func main() {
    nums := []int{2, 3, 1, 4, 2, 3, 1}

    unique := findSingle(nums)

    fmt.Printf("The unique element is: %d\n", unique)
}
