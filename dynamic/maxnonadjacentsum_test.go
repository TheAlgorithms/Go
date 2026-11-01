
package dynamic_test

import (
    "fmt"
    "testing"

    "github.com/TheAlgorithms/Go/dynamic"
)

func TestMaxNonAdjacentSum(t *testing.T) {
    td := []struct {
        input    []int
        expected int
    }{
        {[]int{}, 0},
        {[]int{5}, 5},
        {[]int{-5}, 0},
        {[]int{2, 7, 9, 3, 1}, 12},
        {[]int{1, 2, 3, 4, 5}, 9},
        {[]int{2, 1, 4, 9}, 11},
        {[]int{10, 20}, 20},
        {[]int{-2, -7, -3}, 0},
        {[]int{4, -1, 5, -2, 6}, 15},
        {[]int{5, 5, 10, 100, 10, 5}, 110},
        {[]int{0, 0, 0, 0}, 0},
        {[]int{100, 1, 1, 100}, 200},
    }

    for _, tc := range td {
        name := fmt.Sprintf(
            "MaxNonAdjacentSum(%v)",
            tc.input,
        )

        t.Run(name, func(t *testing.T) {
            actual := dynamic.MaxNonAdjacentSum(tc.input)

            if actual != tc.expected {
                t.Errorf(
                    "MaxNonAdjacentSum(%v) = %d, expected %d",
                    tc.input,
                    actual,
                    tc.expected,
                )
            }
        })
    }
}

func ExampleMaxNonAdjacentSum() {
    fmt.Print(dynamic.MaxNonAdjacentSum([]int{2, 7, 9, 3, 1}))
    // Output: 12
}
