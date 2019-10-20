package bitonic

const (
        SORT_ASC  bool = true
        SORT_DESC bool = false
)

var lst []int

func Sort(input []int, dir bool) {
        lst = input
        sentinel := make(chan int)
        go bitonicSort(0, len(lst), dir, sentinel)
        <-sentinel
}

func bitonicSort(lo int, n int, dir bool, sentinel chan int) {
        if n > 1 {
                m := n / 2
                c1 := make(chan int)
                c2 := make(chan int)
                go bitonicSort(lo, m, SORT_ASC, c1)
                go bitonicSort(lo+m, m, SORT_DESC, c2)
                <-c1
                <-c2
                bitonicMerge(lo, n, dir, sentinel)
        } else {
                sentinel <- 1
        }
}

func bitonicMerge(lo int, n int, dir bool, sentinel chan int) {
        if n > 1 {
                m := n / 2
                for i := lo; i < lo+m; i++ {
                        compareAndSwap(i, i+m, dir)
                }
                c1 := make(chan int)
                c2 := make(chan int)
                go bitonicMerge(lo, m, dir, c1)
                go bitonicMerge(lo+m, m, dir, c2)
                <-c1
                <-c2
        }

        sentinel <- 1
}

func compareAndSwap(i int, j int, dir bool) {
        if dir == (lst[i] > lst[j]) {
                lst[i], lst[j] = lst[j], lst[i]
        }
}