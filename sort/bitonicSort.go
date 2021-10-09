package main

import
(
    "fmt"
    "time"
    "math/rand"
)

const (
    ASC     bool    = true
    DESC    bool    = false
)

var (
    ARR_SIZE    int     = 1<<7
    order       bool    = DESC
)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func bitonic_sort_go(arr []int, orderby bool){
    breakchan := make(chan int)
    go bitonic_sort(arr, breakchan, orderby)
    <-breakchan
}

func bitonic_sort(arr []int ,breakchan chan int, orderby bool){
    if len(arr) < 2 {
        breakchan <- 1
        return
    }

    middle := len(arr) / 2
    c1 := make(chan int)
    c2 := make(chan int)

    go bitonic_sort(arr[:middle], c1, ASC)
    go bitonic_sort(arr[middle:], c2, DESC)
    <-c1
    <-c2
    bitonic_merge(arr, breakchan, orderby)
}

func bitonic_compare(arr []int, orderby bool){
    middle := len(arr) / 2
    for i := 0; i < middle; i++ {
        if (arr[i]>arr[i+middle]) == orderby {
            arr[i], arr[i+middle] = arr[i+middle], arr[i]
        }
    }
}

func bitonic_merge(arr []int, breakchan chan int, orderby bool){
    bitonic_compare(arr, orderby)
    middle := len(arr) / 2
    if middle > 1 {
        c1 := make(chan int)
        c2 := make(chan int)

        go bitonic_merge(arr[:middle], c1, orderby)
        go bitonic_merge(arr[middle:], c2, orderby)
        <-c1
        <-c2
    }
    breakchan <- 1
}

func randInt(min int, max int) int {
    return min+rand.Intn(max-min)
}

func test(orderby bool, arr []int) bool {
    if len(arr) < 2 {
        return true
    }

    prev := arr[0]
    for _, n := range arr[1:]{
        if orderby {
            if prev > n {
                fmt.Println("ASC" ,prev, n)
                return false
            }
        } else
            if prev < n {
                fmt.Println("DESC",prev, n)
                return false
            }
        prev = n
    }
    return true
}

func main(){
    my_arr := make([]int, ARR_SIZE)
    for i := 0; i < ARR_SIZE; i++ {
        my_arr[i] = randInt(1,ARR_SIZE/2)
    }

    bitonic_sort_go(my_arr, order)

    fmt.Println("ARR_SIZE:", ARR_SIZE)
    if order {
        fmt.Println("Order By ASC")
    } else {
        fmt.Println("Order By DESC")
    }
    fmt.Println("my_arr after bitonic sort: ", my_arr)
    fmt.Println("Test: ", test(order, my_arr))
}
