package main

import (
	"fmt"
	array "./arrayList"
	heap "./heap"
)

func isSorted(l *array.List) bool {
	for i:=0; i<l.Length() - 1; i++ {
		if l.GetIndex(i) > l.GetIndex(i+1) {
			return false;
		}
	}

	return true;
}

func bogo(l *array.List){
	for !isSorted(l){
		l.Print()
		l.Shuffle()
	}
	l.Print()
	fmt.Println("Bogo sort works!")
}

func main(){
	var l = array.NewList()
	l.Append(2)
	l.Append(4)
	l.Print()
	for i:=5; i < 12; i++ {
		l.Append(i)
	}
	l.Print()
	l.Swap(0, 1)
	l.Print()
	l.Shuffle()
	l.Print()
	var heap = heap.NewHeap()
	for i := 20; i > 0; i-- {
		heap.Add(i)

	}
	heap.Print()
	for i := 1; i < 21; i++ {
		fmt.Println(heap.RemoveMin())
	}
	fmt.Println("Hello world")
	l.Shuffle()
	var bogod = array.NewList()
	bogod.Append(33)
	bogod.Append(2)
	bogod.Append(15)
	bogo(l)
}