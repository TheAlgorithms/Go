//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"math/rand"
	"sort"
	"time"
)

const (
	maxSize = 500_000
	minSize = -maxSize
)

type sortTest struct {
	input    []int
	expected []int
	name     string
}

var arr []int = makeRandArray()

var sortTests = []sortTest{
	//Sorted slice
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted"},
	//Reversed slice
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed"},
	//Empty slice
	{[]int{}, []int{}, "Empty"},
	//Single-entry slice
	{[]int{1}, []int{1}, "Singleton"},
	// random array of maxSize length sort
	{arr, getSortedVersion(arr), "Large Random"},
}

func makeRandArray() []int {
	rand.Seed(time.Now().UnixNano())

	testArr := []int{}

	for range [maxSize]int{} {
		testArr = append(testArr, rand.Intn(maxSize-minSize)+minSize)
	}
	/* Same method with using crypto rand and only positive numbers
	for i := 0; i < size; i++ {
		temp, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
		vals[i] = int(temp.Int64())
	}
	*/
	return testArr
}

func getSortedVersion(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return a
}
