package graph

func TopoSort(N int, constraints [][]int) ([]int) {
    ans := make([]int, N)
    for i:=0;i<N;i++ {ans[i]=i}
    return ans
}
