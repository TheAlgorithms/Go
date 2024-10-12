package main

import "fmt"

func strandSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    sublist := []int{arr[0]}
    for i := 1; i < len(arr); i++ {
        if arr[i] > sublist[len(sublist)-1] {
            sublist = append(sublist, arr[i])
        }
    }

    remaining := []int{}
    for _, v := range arr {
        if len(sublist) == 0 || v != sublist[0] {
            remaining = append(remaining, v)
        } else {
            sublist = sublist[1:]
        }
    }

    return merge(strandSort(remaining), sublist)
}

func merge(left, right []int) []int {
    result := []int{}
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}

func main() {
    arr := []int{34, 2, 78, 12, 45, 21, 67, 1}
    sorted := strandSort(arr)
    fmt.Println("Strand Sorted array:", sorted)
}
