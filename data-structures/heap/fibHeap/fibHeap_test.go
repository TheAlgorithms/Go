package fibHeap

import (
	"testing"
	"reflect"
	"fmt"
	"sort"
	"math/rand"
)

func TestFibHeapBasic(t *testing.T) {
	arr := []int{1, 3, 2, 2, 4, 5}
	sortedArr := make([]int, 0, 0)
	h := NewFibHeap()
	for _, v := range arr {
		h.Insert(v, v)
	}
	for h.n > 0 {
		min := h.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func TestFibHeap(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	h := NewFibHeap()
	for _, v := range arr {
		h.Insert(v, v)
	}
	for h.n > 0 {
		min := h.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func TestFibHeapUnion(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arrSize1 := rand.Intn(arrSize)
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	h1, h2 := NewFibHeap(), NewFibHeap()
	for i, v := range arr {
		if i < arrSize1 {
			h1.Insert(v, v)
		} else {
			h2.Insert(v, v)
		}
	}
	h1 = h1.Union(h2)
	h2 = nil
	for h1.n > 0 {
		min := h1.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func TestFibHeap_ModifyNode(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	nodeArr := make([]*fibHeapElement, 0, 0)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	h := NewFibHeap()
	for _, v := range arr {
		n := h.Insert(v, v)
		nodeArr = append(nodeArr, n)
	}
	h.consolidate()
	for i := range nodeArr {
		arr[i] += rand.Intn(50) + 50
		h.ModifyNode(nodeArr[i], arr[i], arr[i])
	}
	for h.n > 0 {
		min := h.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func BenchmarkFibHeap_Insert(b *testing.B) {
	b.StopTimer()
	h := NewFibHeap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := rand.Intn(10000)
		b.StartTimer()
		h.Insert(n, n)
	}
}

func BenchmarkFibHeap_ExtractMin(b *testing.B) {
	b.StopTimer()
	h := NewFibHeap()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		h.Insert(n, n)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.ExtractMin()
	}
}

func BenchmarkFibHeap_ModifyNode(b *testing.B) {
	b.StopTimer()
	arr := make([]int, 0, 0)
	nodeArr := make([]*fibHeapElement, 0, 0)
	h := NewFibHeap()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		nodeArr = append(nodeArr,h.Insert(n, n))
		arr = append(arr, n+rand.Intn(50)+50)
	}
	h.consolidate()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.ModifyNode(nodeArr[i], arr[i], arr[i])
	}
}

func BenchmarkFibHeap_ExtractMinAfterModifyNode(b *testing.B) {
	b.StopTimer()
	arr := make([]int, 0, 0)
	nodeArr := make([]*fibHeapElement, 0, 0)
	h := NewFibHeap()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		nodeArr = append(nodeArr,h.Insert(n, n))
		arr = append(arr, n+rand.Intn(50)+50)
	}
	h.consolidate()
	for i := 0; i < b.N; i++ {
		h.ModifyNode(nodeArr[i], arr[i], arr[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.ExtractMin()
	}
}